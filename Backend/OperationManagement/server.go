package main

import (
    "context"
    "math/big"
    "strings"
    "log"
    "fmt"

    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/spf13/viper"
)

func main() {
    // Load configuration from YAML file
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("../..")
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("Fatal error config file: %v \n", err)
    }

    // Read configuration values
    ethereumNodeURL := viper.GetString("ethereumNodeURL")
    privateKeyString := viper.GetString("privateKey")
    contractAddressString := viper.GetString("OperationManagementContractAddress")
    dataCollectionABI := viper.GetString("OperationManagementabi")

    // Connect to Ethereum client
    client, err := ethclient.Dial(ethereumNodeURL)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    // Parse private key
    privateKey, err := crypto.HexToECDSA(privateKeyString)
    if err != nil {
        log.Fatalf("Failed to parse private key: %v", err)
    }

    // Get sender address
    fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

    // Get nonce
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatalf("Failed to get nonce: %v", err)
    }

    // Get gas price
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatalf("Failed to get gas price: %v", err)
    }

    // Parse contract ABI
    contractABI, err := abi.JSON(strings.NewReader(dataCollectionABI))
    if err != nil {
        log.Fatalf("Failed to parse contract ABI: %v", err)
    }

    // Parse contract address
    contractAddress := common.HexToAddress(contractAddressString)

    // Encode transaction data
    actionType := big.NewInt(1) // Change this according to your ActionType enum
    targetDeviceID := big.NewInt(1) // Change this according to your target device ID
    actionParameter := "example" // Change this according to your action parameter
    txData, err := contractABI.Pack("executeAction", actionType, targetDeviceID, actionParameter)
    if err != nil {
        log.Fatalf("Failed to encode transaction data: %v", err)
    }

    // Build transaction
    tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), uint64(3000000), gasPrice, txData)

    // Sign transaction
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), privateKey) // Replace 1 with your chain ID
    if err != nil {
        log.Fatalf("Failed to sign transaction: %v", err)
    }

    // Send transaction
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatalf("Failed to send transaction: %v", err)
    }

    fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
}

