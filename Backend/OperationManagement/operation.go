package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
)

// 定义智能合约的ABI
var abiJSON = `[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "actionType",
				"type": "uint8"
			},
			{
				"indexed": false,
				"name": "timestamp",
				"type": "uint256"
			},
			{
				"indexed": true,
				"name": "executor",
				"type": "address"
			},
			{
				"indexed": false,
				"name": "targetDeviceId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"name": "actionParameter",
				"type": "string"
			}
		],
		"name": "ActionEvent",
		"type": "event"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "actionType",
				"type": "uint8"
			},
			{
				"name": "targetDeviceId",
				"type": "uint256"
			},
			{
				"name": "actionParameter",
				"type": "string"
			}
		],
		"name": "executeAction",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

// 合约地址和私钥（用于发送交易）
const (
	contractAddress = "0x9dC9f7E3c63449803De51ab10CE123504457BbAE"
	privateKey      = "2b955afcaa3a6e5747c02c1f3c5a6ed1295cc351415cbd14affcf73ee979dee8"
)

func main() {
	// 连接到以太坊节点
	client, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 创建一个新的以太坊客户端实例
	ethClient := ethclient.NewClient(client)

	// 解析智能合约ABI
	contractABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		log.Fatal(err)
	}

	// 解析合约地址
	contractAddr := common.HexToAddress(contractAddress)

	// 解析私钥
	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个新的绑定对象
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nil // 如果为nil，将根据网络设置Nonce
	auth.Value = big.NewInt(0) // 使用的Ether数量
	auth.GasLimit = uint64(3000000) // 设置gas限制
	auth.GasPrice = big.NewInt(1000000000) // 设置gas价格

	// 启动模拟
	rand.Seed(time.Now().UnixNano())
	for {
		// 随机选择动作类型、目标设备ID和动作参数
		actionType := rand.Intn(3) // 随机生成0、1、2
		targetDeviceId := rand.Uint64()
		actionParameter := randomString(6)

		// 调用智能合约的executeAction函数
		tx, err := contractABI.Pack("executeAction", uint8(actionType), targetDeviceId, actionParameter)
		if err != nil {
			log.Fatal(err)
		}

		// 发送交易
		txHash, err := ethClient.SendTransaction(auth, tx, &contractAddr)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Action executed: ActionType=%d, TargetDeviceId=%d, ActionParameter=%s, TxHash=%s\n", actionType, targetDeviceId, actionParameter, txHash.Hex())

		// 休眠一段时间，模拟人为操作的时间间隔
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

// 生成随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
