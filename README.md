# Ethereum Go Client

This is a Ethereum compatible Go Client
which implements the 
[Eth JSON RPC Module](https://github.com/ethereum/wiki/wiki/JSON-RPC),
[Personal JSON RPC Module](https://github.com/paritytech/parity/wiki/JSONRPC-personal-module) and
[NET JSON RPC Module](https://github.com/paritytech/parity/wiki/JSONRPC-net-module#net_version).

## Status

This package is currently under active development. It is not already stable and the infrastructure is not complete and there are still several RPCs left to implement and the API is not stable yet.


## Installation

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

## Usage

```go
import (
	web3 "github.com/fraymond/web3go"
	"github.com/fraymond/web3go/eth/block"
	"github.com/fraymond/web3go/providers"
	"fmt"
)

func test() {

	web3Client := web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:30305", 10, true))
	balance, err := web3Client.Eth.GetBalance("0x24fc5c1c97f838e57c944240fa2ffc18256bc415", block.LATEST)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(balance)

}
```

More tests in the 'tests' directory.

### Requirements

* go ^1.8.3

[Go installation instructions.](https://golang.org/doc/install)


## Testing
```bash
go test -v test/*.go
```

## License

Package web3go is licensed under the [GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html) License.
