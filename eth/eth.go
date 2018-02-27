/********************************************************************************
   This file is part of web3go.
   web3go is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   web3go is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with web3go.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file eth.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */

package eth

import (
	"github.com/fraymond/web3go/complex/types"
	"github.com/fraymond/web3go/dto"
	"github.com/fraymond/web3go/providers"
)

// Eth - The Eth Module
type Eth struct {
	provider providers.ProviderInterface
}

// NewEth - Eth Module constructor to set the default provider
func NewEth(provider providers.ProviderInterface) *Eth {
	eth := new(Eth)
	eth.provider = provider
	return eth
}

// GetProtocolVersion - Returns the current ethereum protocol version.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_protocolversion
// Parameters:
//    - none
// Returns:
// 	  - String - The current ethereum protocol version
func (eth *Eth) GetProtocolVersion() (string, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_protocolVersion", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// IsSyncing - Returns an object with data about the sync status or false.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_syncing
// Parameters:
//    - none
// Returns:
// 	  - Object|Boolean, An object with sync status data or FALSE, when not syncing:
//    	- startingBlock: 	QUANTITY - The block at which the import started (will only be reset, after the sync reached his head)
//    	- currentBlock: 	QUANTITY - The current block, same as eth_blockNumber
//    	- highestBlock: 	QUANTITY - The estimated highest block
func (eth *Eth) IsSyncing() (*dto.SyncingResponse, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_syncing", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToSyncingResponse()

}

// GetCoinbase - Returns the client coinbase address.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_coinbase
// Parameters:
//    - none
// Returns:
// 	  - DATA, 20 bytes - the current coinbase address.
func (eth *Eth) GetCoinbase() (string, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_coinbase", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// IsMining - Returns true if client is actively mining new blocks.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_mining
// Parameters:
//    - none
// Returns:
// 	  - Boolean - returns true of the client is mining, otherwise false.
func (eth *Eth) IsMining() (bool, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_mining", nil)

	if err != nil {
		return false, err
	}

	return pointer.ToBoolean()

}

// GetHashRate - Returns the number of hashes per second that the node is mining with.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_hashrate
// Parameters:
//    - none
// Returns:
// 	  - QUANTITY - number of hashes per second.
func (eth *Eth) GetHashRate() (types.ComplexIntResponse, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_hashrate", nil)

	if err != nil {
		return types.ComplexIntResponse(0), err
	}

	return pointer.ToComplexIntResponse()

}

// GetGasPrice - Returns the current price per gas in wei.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gasprice
// Parameters:
//    - none
// Returns:
// 	  - QUANTITY - integer of the current gas price in wei.
func (eth *Eth) GetGasPrice() (types.ComplexIntResponse, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_gasPrice", nil)

	if err != nil {
		return types.ComplexIntResponse(0), err
	}

	return pointer.ToComplexIntResponse()

}

// ListAccounts - Returns a list of addresses owned by client.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_accounts
// Parameters:
//    - none
// Returns:
//    - Array of DATA, 20 Bytes - addresses owned by the client.
func (eth *Eth) ListAccounts() ([]string, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_accounts", nil)

	if err != nil {
		return nil, err
	}

	return pointer.ToStringArray()

}

// GetBlockNumber - Returns the number of most recent block.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_blocknumber
// Parameters:
//    - none
// Returns:
// 	  - QUANTITY - integer of the current block number the client is on.
func (eth *Eth) GetBlockNumber() (types.ComplexIntResponse, error) {

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_blockNumber", nil)

	if err != nil {
		return types.ComplexIntResponse(0), err
	}

	return pointer.ToComplexIntResponse()

}

// GetBalance - Returns the balance of the account of given address.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getbalance
// Parameters:
//    - DATA, 20 Bytes - address to check for balance.
//	  - QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter: https://github.com/ethereum/wiki/wiki/JSON-RPC#the-default-block-parameter
// Returns:
// 	  - QUANTITY - integer of the current balance in wei.
func (eth *Eth) GetBalance(address string, defaultBlockParameter string) (types.ComplexIntResponse, error) {

	params := make([]string, 2)
	params[0] = address
	params[1] = defaultBlockParameter

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_getBalance", params)

	if err != nil {
		return "", err
	}

	return pointer.ToComplexIntResponse()

}

// GetStorageAt - Returns the value from a storage position at a given address.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getstorageat
// Parameters:
//    - DATA, 20 Bytes - address of the storage.
//	  - QUANTITY - integer of the position in the storage.
//	  - QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter: https://github.com/ethereum/wiki/wiki/JSON-RPC#the-default-block-parameter.
// Returns:
// 	  - DATA - the value at this storage position.
func (eth *Eth) GetStorageAt(address string, position types.ComplexIntParameter, defaultBlockParameter string) (string, error) {

	params := make([]string, 3)
	params[0] = address
	params[1] = position.ToHex()
	params[2] = defaultBlockParameter

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_getstorageat", params)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// EstimateGas - Makes a call or transaction, which won't be added to the blockchain and returns the used gas, which can be used for estimating the used gas.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_estimategas
// Parameters:
//    - See eth_call parameters, expect that all properties are optional. If no gas limit is specified geth uses the block gas limit from the pending block as an
// 		upper bound. As a result the returned estimate might not be enough to executed the call/transaction when the amount of gas is higher than the pending block gas limit.
// Returns:
//    - QUANTITY - the amount of gas used.
func (eth *Eth) EstimateGas(transaction *dto.TransactionParameters) (types.ComplexIntResponse, error) {

	params := make([]*dto.RequestTransactionParameters, 1)

	params[0] = transaction.Transform()

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(&pointer, "eth_estimateGas", params)

	if err != nil {
		return types.ComplexIntResponse(0), err
	}

	return pointer.ToComplexIntResponse()

}

// GetTransactionByHash - Returns the information about a transaction requested by transaction hash.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionbyhash
// Parameters:
//    - DATA, 32 Bytes - hash of a transaction
// Returns:
//    1. Object - A transaction object, or null when no transaction was found
//    - hash: DATA, 32 Bytes - hash of the transaction.
//    - nonce: QUANTITY - the number of transactions made by the sender prior to this one.
//    - blockHash: DATA, 32 Bytes - hash of the block where this transaction was in. null when its pending.
//    - blockNumber: QUANTITY - block number where this transaction was in. null when its pending.
//    - transactionIndex: QUANTITY - integer of the transactions index position in the block. null when its pending.
//    - from: DATA, 20 Bytes - address of the sender.
//    - to: DATA, 20 Bytes - address of the receiver. null when its a contract creation transaction.
//    - value: QUANTITY - value transferred in Wei.
//    - gasPrice: QUANTITY - gas price provided by the sender in Wei.
//    - gas: QUANTITY - gas provided by the sender.
//    - input: DATA - the data send along with the transaction.
func (eth *Eth) GetTransactionByHash(hash string) (*dto.TransactionResponse, error) {

	params := make([]string, 1)
	params[0] = hash

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_getTransactionByHash", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToTransactionResponse()

}

// SendTransaction - Creates new message call transaction or a contract creation, if the data field contains code.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendtransaction
// Parameters:
//    1. Object - The transaction object
//    - from: 		DATA, 20 Bytes - The address the transaction is send from.
//    - to: 		DATA, 20 Bytes - (optional when creating new contract) The address the transaction is directed to.
//    - gas: 		QUANTITY - (optional, default: 90000) Integer of the gas provided for the transaction execution. It will return unused gas.
//    - gasPrice: 	QUANTITY - (optional, default: To-Be-Determined) Integer of the gasPrice used for each paid gas
//    - value: 		QUANTITY - (optional) Integer of the value send with this transaction
//    - data: 		DATA - The compiled code of a contract OR the hash of the invoked method signature and encoded parameters. For details see Ethereum Contract ABI (https://github.com/ethereum/wiki/wiki/Ethereum-Contract-ABI)
//    - nonce: 		QUANTITY - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.
// Returns:
//	  - DATA, 32 Bytes - the transaction hash, or the zero hash if the transaction is not yet available.
// Use eth_getTransactionReceipt to get the contract address, after the transaction was mined, when you created a contract.
func (eth *Eth) SendTransaction(transaction *dto.TransactionParameters) (string, error) {

	params := make([]*dto.RequestTransactionParameters, 1)
	params[0] = transaction.Transform()

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(&pointer, "eth_sendTransaction", params)

	if err != nil {
		return "", err
	}

	response, err := pointer.ToString()

	return response, err

}

// CompileSolidity - Returns compiled solidity code.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_compilesolidity
// Parameters:
//    1. String - The source code.
// Returns:
//	  - DATA - The compiled source code.
func (eth *Eth) CompileSolidity(sourceCode string) (types.ComplexString, error) {

	params := make([]string, 1)
	params[0] = sourceCode

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_compileSolidity", params)

	if err != nil {
		return "", err
	}

	return pointer.ToComplexString()

}

// GetTransactionReceipt - Returns compiled solidity code.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionreceipt
// Parameters:
//    1. DATA, 32 Bytes - hash of a transaction.
// Returns:
//	  1. Object - A transaction receipt object, or null when no receipt was found:
//    - transactionHash: 		DATA, 32 Bytes - hash of the transaction.
//    - transactionIndex: 		QUANTITY - integer of the transactions index position in the block.
//    - blockHash: 				DATA, 32 Bytes - hash of the block where this transaction was in.
//    - blockNumber:			QUANTITY - block number where this transaction was in.
//    - cumulativeGasUsed: 		QUANTITY - The total amount of gas used when this transaction was executed in the block.
//    - gasUsed: 				QUANTITY - The amount of gas used by this specific transaction alone.
//    - contractAddress: 		DATA, 20 Bytes - The contract address created, if the transaction was a contract creation, otherwise null.
//    - logs: 					Array - Array of log objects, which this transaction generated.
func (eth *Eth) GetTransactionReceipt(hash string) (*dto.TransactionReceipt, error) {

	params := make([]string, 1)
	params[0] = hash

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_getTransactionReceipt", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToTransactionReceipt()

}

// GetBlockByNumber - Returns the information about a block requested by number.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbynumber
// Parameters:
//    - number, QUANTITY - number of block
//    - transactionDetails, bool - indicate if we should have or not the details of the transactions of the block
// Returns:
//    1. Object - A block object, or null when no transaction was found
//    2. error
func (eth *Eth) GetBlockByNumber(number types.ComplexIntParameter, transactionDetails bool) (*dto.Block, error) {

	params := make([]interface{}, 2)
	params[0] = number.ToHex()
	params[1] = transactionDetails

	pointer := &dto.RequestResult{}

	err := eth.provider.SendRequest(pointer, "eth_getBlockByNumber", params)

	if err != nil {
		return nil, err
	}

	return pointer.ToBlock()

}
