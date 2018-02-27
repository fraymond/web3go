# Ethereum Go Client

This is a Ethereum compatible Go Client
which implements the 
[Eth JSON RPC Module](https://github.com/ethereum/wiki/wiki/JSON-RPC). Similar to [Web3js](https://github.com/ethereum/web3.js/).


## Web3go Installation

### go get

```bash
go get -u github.com/fraymond/web3go
```

## Ethereum Configuration

### Install Ethereum
```
cd ~/git
git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
make geth
export PATH=$PATH:~/git/go-ethereum/build/bin/
```

### Run Ethereum

```
mkdir ~/geth-test
cd ~/geth-test
mkdir privatebc
```
create genesis.json file

```
vi genesis.json
```
put the following in genesis.json file and save.
```
{
    "config": {  
        "chainId": 4321, 
        "homesteadBlock": 0,
        "eip155Block": 0,
        "eip158Block": 0
    },
    "difficulty": "0x400",
    "gasLimit": "0x8000000",
    "alloc": {}
}
```
Start geth and open a web3js console. You can use port 30305, which is default for web3go testing.

```
geth --datadir ./privatebc init ./genesis.json
geth --datadir ./privatebc --networkid 4321 --port 30305 console
```

Create new accounts and send transactions

```
eth.coinbase
eth.accounts
personal.newAccount()
passphrase:
repeat passphrase:

miner.start()
--wait a few seconds
miner.stop()

personal.unlockAccount("0x18833df6ba69b4d50acc744e8294d128ed8db1f1")
eth.sendTransaction({from: '0x18833df6ba69b4d50acc744e8294d128ed8db1f1', to: '0x2a022eb956d1962d867dcebd8fed6ae71ee4385a', value: web3.toWei(12, "ether")}) 
```

## Web3go Execution
```bash
go run web3main.go
```

### Requirements

* go ^1.8.3

[Go installation instructions.](https://golang.org/doc/install)
