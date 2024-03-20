package main

import (
    
    "log"
    "myproject/BlockChain"

)

// 假设SendHashToContract函数的实现已经在此处或其他文件中

func main() {
    exampleHash := "0e5751c026e543b2e8ab2c1a1f3a3f5a3a38e2d5df3345ad284ed481edb077f0"
    exampleDeviceId := int64(123) // 模拟的设备ID

    err := BlockChain.SendHashToContract(exampleDeviceId, exampleHash)
    if err != nil {
        log.Fatalf("Error sending hash to contract: %v\n", err)
    }
}

