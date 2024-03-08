# Geth Installation

It's recommended to use apt for installation to avoid the complex compilation process, which is prone to errors. [Geth Installation Guide](https://geth.ethereum.org/docs/getting-started/installing-geth#install-on-ubuntu-via-ppas) provides a straightforward method for Ubuntu-based distributions using the built-in launchpad PPAs (Personal Package Archives). A single PPA repository contains both stable and development releases for Ubuntu versions xenial, trusty, impish, focal, bionic.

To enable the launchpad repository, use the command:

```sh
sudo add-apt-repository -y ppa:ethereum/ethereum
```

To install the stable version of go-ethereum, run:

```sh
sudo apt-get update
sudo apt-get install ethereum
```

Or, for the develop version:

```sh
sudo apt-get update
sudo apt-get install ethereum-unstable
```

This installation includes the core Geth software along with developer tools such as clef, devp2p, abigen, bootnode, evm, and rlpdump, which are saved in /usr/local/bin/. The full list of command-line options is available [here](https://geth.ethereum.org/docs/fundamentals/Command-Line-Options) or by running `geth --help` in the terminal.

To update an existing Geth installation, stop the node and run:

```sh
sudo apt-get update
sudo apt-get install ethereum
sudo apt-get upgrade geth
```

Upon restarting, Geth will automatically utilize all data from the previous version and sync any missed blocks.

# Ethereum Initialization Under Ethash Consensus

## Creating an Account

To create an account, use:

```sh
geth account new --datadir data
```

Enter a password when prompted and note the public key (pk) address. For encoding the signer's address in extradata, append 32 zero bytes, all signer addresses, and another 65 zero bytes.

## Modifying the Initial Configuration File (genesis.json)

Edit `genesis.json` with the specific configuration for your network, including `chainId`, consensus parameters, and `extradata` with the encoded addresses.

## Initialization

To initialize the network with your genesis file, run:

```sh
geth init --datadir data genesis.json
```

## Starting the Node

### Test Start

```sh
geth --datadir data --networkid 12345
```

### Starting the Chain, Console, and Redirecting Log Output

```sh
geth --datadir data --datadir data --networkid 12345 --mine --miner.etherbase=<YourAccountAddress> --rpc.enabledeprecatedpersonal --allow-insecure-unlock  console 2>geth.log
```

Replace `<YourAccountAddress>` with the address from your account creation step. This command starts mining and opens a console for interaction, with logs directed to `geth.log`.

### Dynamic Tracking

```sh
tail -f geth.log
```

This allows for monitoring the log output in real time without affecting command input in the console.

# Web3.js Interaction

## Creating an Account

If encountering errors like “ReferenceError: personal is not defined,” restart the Geth node with the `--rpc.enabledeprecatedpersonal` flag during node startup. This is necessary for Geth versions 1.11.* and above, where the personal API has been removed. Earlier versions do not require this flag.

```sh
personal.newAccount()
```

## Querying Accounts

To list accounts and check balances:

```shell
eth.accounts
eth.getBalance(eth.accounts[0])
```

## Mining

Unlock the account for mining and start or stop the miner as needed:

```
personal.unlockAccount(eth.accounts[0], "password", 0)
miner.start()
miner.stop()
```

To check the block number and hash rate:

```
eth.blockNumber
eth.hashrate
```

Note: The `miner.hashrate` command is only applicable to CPU mining. For GPU mining, it always reports zero.
