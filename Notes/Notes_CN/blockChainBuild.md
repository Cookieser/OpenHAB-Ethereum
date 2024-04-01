# Geth安装

建议直接使用apt进行安装，不需要进行复杂的编译过程（还容易出错）
[https://geth.ethereum.org/docs/getting-started/installing-geth#install-on-ubuntu-via-ppas](https://geth.ethereum.org/docs/getting-started/installing-geth#install-on-ubuntu-via-ppas)

The easiest way to install Geth on Ubuntu-based distributions is with the built-in launchpad PPAs (Personal Package Archives). A single PPA repository is provided, containing stable and development releases for Ubuntu versions xenial, trusty, impish, focal, bionic.

The following command enables the launchpad repository:

```sh
sudo add-apt-repository -y ppa:ethereum/ethereum
```

Then, to install the stable version of go-ethereum:

```sh
sudo apt-get update
sudo apt-get install ethereum
```

Or, alternatively the develop version:

```sh
sudo apt-get update
sudo apt-get install ethereum-unstable
```

These commands install the core Geth software and the following developer tools: clef, devp2p, abigen, bootnode, evm and rlpdump. The binaries for each of these tools are saved in /usr/local/bin/. The full list of command line options can be viewed [here](https://geth.ethereum.org/docs/fundamentals/Command-Line-Options) or in the terminal by running geth --help.

Updating an existing Geth installation to the latest version can be achieved by stopping the node and running the following commands:

```sh
sudo apt-get update
sudo apt-get install ethereum
sudo apt-get upgrade geth
```

When the node is started again, Geth will automatically use all the data from the previous version and sync the blocks that were missed while the node was offline.



# Ethash共识下的以太坊初始化

[https://geth.ethereum.org/docs/fundamentals/private-network](https://geth.ethereum.org/docs/fundamentals/private-network)、

## 创建账号

```sh
geth account new --datadir data
password: test
pk: 0x13524237992D83974367E5fb1B29f6Fcbe1ba3Df
```

记录pk地址。要在extradata中对签名者地址进行编码，请连接 32 个零字节、所有签名者地址和另外 65 个零字节。

## 修改初始配置文件genesis.json

```
{
  "config": {
    "chainId": 12345,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "berlinBlock": 0,
    "clique": {
      "period": 5,
      "epoch": 30000
    }
  },
  "difficulty": "1",
  "gasLimit": "8000000",
  "extradata": "0x000000000000000000000000000000000000000000000000000000000000000013524237992D83974367E5fb1B29f6Fcbe1ba3Df0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  "alloc": {
    "7df9a875a174b3bc565e6424a0050ebc1b2d1d82": { "balance": "300000" },
    "f41c74c9ae680c1aa78f42e5647a62f353b7bdde": { "balance": "400000" }
  }
}
```

## 初始化

```sh
geth init --datadir data genesis.json
```

## 启动

### 测试启动

```sh
geth --datadir data --networkid 12345
```

### 启动链、控制台并重定向日志输出

```sh
geth --datadir data --datadir data --networkid 12345 --mine --miner.etherbase=0x6751E595E7B7a728C3E41bD0c19C4d5219b49d71 --rpc.enabledeprecatedpersonal --allow-insecure-unlock  console 2>geth.log
# 注意根据上述的account修改miner.etherbase参数
```

### 动态跟踪

```sh
tail -f geth.log
```

可以分开控制台和输出，避免输出影响命令输入



# Web3.js交互

## 创建账户

> 如果此命令不起作用（通常会出现诸如“ReferenceError：personal is not Defined”之类的错误），那么您必须在启动节点时（而不是附加时）使用该--rpc.enabledeprecatedpersonal 标志重新启动geth节点。这种情况发生在 geth 1.11.* 版本上，该版本删除了个人 API；版本 1.10.* 及更低版本不需要此标志。

```shell
personal.newAccount()
```

查询账户

```
eth.accounts
```

```
eth.getBalance(eth.accounts[0])
```

## 挖矿

```
personal.unlockAccount(eth.accounts[0], "test", 0)
```

```
miner.start()
miner.stop()
```

```
eth.blockNumber
```

```
eth.hashrate
```

注意：Geth 命令 miner.hashrate 仅适用于 CPU 挖掘 - 对于 GPU 挖掘，它始终报告零

