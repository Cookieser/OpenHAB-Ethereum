package main

import (
    "context"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {

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
