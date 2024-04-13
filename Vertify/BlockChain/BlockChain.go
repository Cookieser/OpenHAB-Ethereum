package BlockChain

import (
    "context"
    "encoding/hex"
    "fmt"
    "log"
    "math/big"
    "strings"

    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/spf13/viper"
)

// TemperatureData represents the structure for storing temperature data.
type TemperatureData struct {
    DeviceId    int64   `json:"deviceId"`
    Temperature float64 `json:"temperature"`
    Unit        string  `json:"unit"`
    Timestamp   string  `json:"timestamp"`
}

// Init initializes a connection to the Ethereum blockchain.
func Init() {
    client, err := ethclient.Dial("ws://localhost:7545")
    if err != nil {
        log.Fatalf("Unable to connect to Ethereum client: %v", err)
    }

    blockNumber, err := client.BlockNumber(context.Background())
    if err != nil {
        log.Fatalf("Failed to get the latest block number: %v", err)
    }
    fmt.Printf("Latest block number: %d\n", blockNumber)

    fmt.Println("Connection to Ethereum network successful!")
}

// SendHashToContract sends a hash to the smart contract on the blockchain.
func SendHashToContract(deviceId int64, dataHashString string) error {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("..")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Fatal error reading config file: %v \n", err)
    }

    ethereumNodeURL := viper.GetString("ethereumNodeURL")
    privateKeyString := viper.GetString("privateKey")
    contractAddressString := viper.GetString("DataCollectionContractAddress")
    dataCollectionABI := viper.GetString("DataCollectionabi")

    privateKey, err := crypto.HexToECDSA(privateKeyString)
    if err != nil {
        return fmt.Errorf("unable to parse private key: %v", err)
    }
    client, err := ethclient.Dial(ethereumNodeURL)
    if err != nil {
        return fmt.Errorf("unable to connect to Ethereum client: %v", err)
    }

    fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        return fmt.Errorf("unable to get pending nonce: %v", err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        return fmt.Errorf("unable to suggest gas price: %v", err)
    }

    contractABI, err := abi.JSON(strings.NewReader(dataCollectionABI))
    if err != nil {
        return fmt.Errorf("unable to parse ABI: %v", err)
    }

    contractAddress := common.HexToAddress(contractAddressString)

    dataHash, err := hex.DecodeString(dataHashString)
    if err != nil {
        return fmt.Errorf("unable to decode hash string: %v", err)
    }

    if len(dataHash) != 32 {
        return fmt.Errorf("hashed data is not the correct length to be a bytes32")
    }

    var dataHashArray [32]byte
    copy(dataHashArray[:], dataHash)

    txData, err := contractABI.Pack("recordDataHash", deviceId, dataHashArray)
    if err != nil {
        return fmt.Errorf("unable to pack transaction data: %v", err)
    }

    chainID := big.NewInt(1337) // Replace with your chain ID
    tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), uint64(3000000), gasPrice, txData)
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        return fmt.Errorf("unable to sign transaction: %v", err)
    }

    if err := client.SendTransaction(context.Background(), signedTx); err != nil {
        return fmt.Errorf("unable to send transaction: %v", err)
    }

    fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
    return nil
}

