package main

import (
    "context"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "strings"
)

func main() {
    client, err := ethclient.Dial("ws://localhost:7545")
    if err != nil {
        log.Fatalf("Wrong: %v", err)
    }
    fmt.Println("Success!!!")

    contractAddress := common.HexToAddress("0xdB7Ff34cb23d1d2e6CC8e8f87b8f7565eF36a8Bd")
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    contractAbi, err := abi.JSON(strings.NewReader(`[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "_fName",
				"type": "string"
			},
			{
				"indexed": true,
				"name": "_age",
				"type": "uint256"
			}
		],
		"name": "InfoSet",
		"type": "event"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_fName",
				"type": "string"
			},
			{
				"name": "_age",
				"type": "uint256"
			}
		],
		"name": "setInfo",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "age",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "fName",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]`))
    if err != nil {
        log.Fatal(err)
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Listening for InfoSet events...")
    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            fmt.Println("Event detected:", vLog)

            // Assuming Age is indexed and FName is not
            var eventName = "InfoSet"
            event := struct {
                FName string
                Age   uint
            }{}
            err := contractAbi.UnpackIntoInterface(&event, eventName, vLog.Data)
            if err != nil {
                log.Fatalf("Failed to unpack: %v", err)
            }

            // Age is indexed, so it's stored in Topics[1] not in Data
            // Assuming Age is the first and only indexed parameter
            var age = new(big.Int)
            if len(vLog.Topics) > 1 {
                age.SetBytes(vLog.Topics[1].Bytes())
            }

            fmt.Printf("InfoSet event detected, FName: %s, Age: %d\n", event.FName, age.Uint64())
        }
    }
}
