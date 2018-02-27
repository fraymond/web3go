package main

import (
	"fmt"

	web3 "github.com/fraymond/web3go"
	"github.com/fraymond/web3go/providers"
	"github.com/fraymond/web3go/dto"
)

func main(){

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:30305", 10, false))

	accounts, err := connection.Personal.ListAccounts()

	if err != nil {
		fmt.Println(err)
	}

	result, err := connection.Personal.UnlockAccount(accounts[0], "0000", 100)

	if err != nil {
		fmt.Println(err)
	}

	if !result {
		fmt.Println("Can't unlock account")
	}

	transaction := new(dto.TransactionParameters)
	transaction.Data = "test"
	transaction.From = accounts[0]
	transaction.To = accounts[1]
	transaction.Value = 10
	transaction.Gas = 40000

	txID, err := connection.Personal.SendTransaction(transaction, "0000")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(txID)


}