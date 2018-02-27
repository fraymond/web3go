package main

import (
	"fmt"

	web3 "github.com/fraymond/web3go"
	"github.com/fraymond/web3go/providers"
	"github.com/fraymond/web3go/dto"
)
/**
 * @file web3main.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 * The main function will connect to an ethereum network, get a list of accounts.
 * unlock the first account and send transaction to a second account.
 */
func main(){

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 100, false))

	coinbase, err := connection.Eth.GetCoinbase()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(coinbase)

	accounts, err := connection.Personal.ListAccounts()

	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := connection.Personal.UnlockAccount(accounts[0], "0000", 100)

	if err != nil {
		fmt.Println(err)
		return
	}

	if !result {
		fmt.Println("Can't unlock account")
		return
	}

	transaction := new(dto.TransactionParameters)
	transaction.Data = "test"
	transaction.From = accounts[0] //"0x18833df6ba69b4d50acc744e8294d128ed8db1f1" //accounts[0]
	transaction.To = accounts[1] //"0x882dbeb3de07f01df95e14e9db16d834a8ceea8f" //accounts[1]
	transaction.Value = 10
	transaction.Gas = 40000

	txID, err := connection.Eth.SendTransaction(transaction)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(txID)


}