package BlockChain

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "github.com/spf13/viper"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
   "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "strings"
    "encoding/hex"


)

type TemperatureData struct {
    DeviceId    int64  `json:"deviceId"`
    Temperature float64 `json:"temperature"`
    Unit        string  `json:"unit"`
    Timestamp   string  `json:"timestamp"`
}

func Init() {

    client, err := ethclient.Dial("http://localhost:7545")
    if err != nil {
        log.Fatalf("Wrong: %v", err)
    }
    fmt.Println("Connection to BlockChain -------------------------------- Success!!!")
    
    
    
    blockNumber, err := client.BlockNumber(context.Background())
    if err != nil {
        log.Fatalf("Failed to get the latest block number: %v", err)
    }
    fmt.Printf("Latest block number: %d\n", blockNumber)

    fmt.Println("Connection to Ethereum network successful!")
}


/*
func SendToContract(data TemperatureData) error {

    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
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
}*/

// Adjust the function signature to accept device ID and hashed data
func SendHashToContract(deviceId int64, dataHashString string) error {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("..")
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





// Convert the string hash back to a byte slice to match the expected type
dataHash, err := hex.DecodeString(dataHashString)
if err != nil {
    return err
}

// Ensure the byte slice is the correct length for a bytes32 type
if len(dataHash) != 32 {
    return fmt.Errorf("hashed data is not the correct length to be a bytes32")
}

// Convert slice to array
var dataHashArray [32]byte
copy(dataHashArray[:], dataHash)

// Create transaction data using the `recordDataHash` method, adjusting for the correct parameter types
txData, err := contractABI.Pack("recordDataHash", deviceId, dataHashArray) // 使用数组而不是切片
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


