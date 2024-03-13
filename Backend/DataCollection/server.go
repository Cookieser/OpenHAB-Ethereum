package main

import (
    "context"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/big"
    "net/http"
    "strings"
    "time"
    "log"

    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/spf13/viper"
)

type TemperatureData struct {
    DeviceId    int64  `json:"deviceId"`
    Temperature float64 `json:"temperature"`
    Unit        string  `json:"unit"`
    Timestamp   string  `json:"timestamp"`
}

func main() {

    

    http.HandleFunc("/temperature", handleTemperature)
    fmt.Println("Server is listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}

func handleTemperature(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()

    var data TemperatureData
    if err := json.Unmarshal(body, &data); err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }

    hashedData, err := hashTemperatureData(data)
    if err != nil {
        fmt.Println("Error hashing data:", err)
        http.Error(w, "Error hashing data", http.StatusInternalServerError)
        return
    }

    fmt.Printf("Received data: %+v\n", data)
    fmt.Println("Hashed Data:", hashedData)


    if err := sendToContract(data); err != nil {
        http.Error(w, "Error sending data to the contract", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Data received and processed"))
}

func hashTemperatureData(data TemperatureData) (string, error) {
    dataString := fmt.Sprintf("%d:%f:%s", data.DeviceId, data.Temperature, data.Timestamp)
    hash := sha256.New()
    _, err := hash.Write([]byte(dataString))
    if err != nil {
        return "", err
    }
    hashedData := hex.EncodeToString(hash.Sum(nil))
    return hashedData, nil
}

func sendToContract(data TemperatureData) error {

    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("../..")
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("Fatal error config file: %v \n", err)
    }

    ethereumNodeURL := viper.GetString("ethereumNodeURL")
    privateKeyString := viper.GetString("privateKey")
    contractAddressString := viper.GetString("DataCollectionContractAddress")
    dataCollectionABI := viper.GetString("DataCollectionabi")





    client, err := ethclient.Dial(ethereumNodeURL)
    if err != nil {
        return err
    }

    privateKey, err := crypto.HexToECDSA(privateKeyString)
    if err != nil {
        return err
    }

    fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        return err
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        return err
    }

    contractABI, err := abi.JSON(strings.NewReader(dataCollectionABI))
    if err != nil {
        return err
    }

    contractAddress := common.HexToAddress(contractAddressString)
    timestamp, _ := time.Parse(time.RFC3339, data.Timestamp)
    txData, err := contractABI.Pack("addData", big.NewInt(int64(data.DeviceId)), big.NewInt(timestamp.Unix()), big.NewInt(int64(data.Temperature)))
    if err != nil {
        return err
    }

    chainID := big.NewInt(1337) // Replace with your chain ID
    tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), uint64(3000000), gasPrice, txData)
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        return err
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return err
    }

    fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
    return nil
}
