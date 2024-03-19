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
    "time"
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
        log.Fatalf("无法连接到以太坊客户端: %v", err)
    }
    fmt.Println("成功连接到以太坊客户端")


    account := common.HexToAddress("0xB4ba642eF0C62aF4ED6fc8cfc2d09001232884Ec")
    balance, err := client.BalanceAt(context.Background(), account, nil) 
    if err != nil {
        log.Fatalf("无法获取账户余额: %v", err)
    }
    fmt.Println("成功获取账户余额")

    balanceInEther := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
    fmt.Printf("账户余额: %f ETH\n", balanceInEther)
}



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
}
