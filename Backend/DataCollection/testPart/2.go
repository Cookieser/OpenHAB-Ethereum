package main

import (
    "context"
    
    "fmt"
    "log"
    "math/big"
    "strings"
    "time"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/spf13/viper"
)

func main() {
   
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("..") //has been moved!!!
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
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    
    privateKey, err := crypto.HexToECDSA(privateKeyString)
    if err != nil {
        log.Fatal(err)
    }

    fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

   
    contractABI, err := abi.JSON(strings.NewReader(dataCollectionABI))
    if err != nil {
        log.Fatal(err)
    }

    
    contractAddress := common.HexToAddress(contractAddressString)

    
    data, err := contractABI.Pack("addData", big.NewInt(123), big.NewInt(time.Now().Unix()), big.NewInt(456))
    if err != nil {
        log.Fatal(err)
    }

    chainID := big.NewInt(1337) 
    tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), uint64(300000), gasPrice, data)

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }

    
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
}

