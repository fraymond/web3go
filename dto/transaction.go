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
 * @file transaction.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */

package dto

import (
	"github.com/fraymond/web3go/complex/types"
)

// TransactionParameters GO transaction to make more easy controll the parameters
type TransactionParameters struct {
	From     string
	To       string
	Gas      types.ComplexIntParameter
	GasPrice types.ComplexIntParameter
	Value    types.ComplexIntParameter
	Data     types.ComplexString
}

// RequestTransactionParameters JSON
type RequestTransactionParameters struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Value    string `json:"value"`
	Data     string `json:"data,omitempty"`
}

// Transform the GO transactions parameters to json style
func (params *TransactionParameters) Transform() *RequestTransactionParameters {
	request := new(RequestTransactionParameters)
	request.From = params.From
	request.To = params.To
	if params.Gas != 0 {
		request.Gas = params.Gas.ToHex()
	} else {
		request.Gas = "0x0"
	}
	if params.GasPrice != 0 {
		request.GasPrice = params.GasPrice.ToHex()
	} else {
		request.GasPrice = "0x0"
	}
	if params.Value != 0 {
		request.Value = params.Value.ToHex()
	} else {
		request.Value = "0x0"
	}
	if params.Data != "" {
		request.Data = params.Data.ToHex()
	} else {
		request.Data = "0x0"
	}
	return request
}

type TransactionResponse struct {
	Hash             string                   `json:"hash"`
	Nonce            int                      `json:"nonce"`
	BlockHash        string                   `json:"blockHash"`
	BlockNumber      int64                    `json:"blockNumber"`
	TransactionIndex int64                    `json:"transactionIndex"`
	From             string                   `json:"from"`
	To               string                   `json:"to"`
	Value            types.ComplexIntResponse `json:"value"`
	GasPrice         types.ComplexIntResponse `json:"gasPrice,omitempty"`
	Gas              types.ComplexIntResponse `json:"gas,omitempty"`
	Data             types.ComplexString      `json:"data,omitempty"`
}

type TransactionReceipt struct {
	TransactionHash   string   `json:"transactionHash"`
	TransactionIndex  int64    `json:"transactionIndex"`
	BlockHash         string   `json:"blockHash"`
	BlockNumber       int64    `json:"blockNumber"`
	CumulativeGasUsed int64    `json:"cumulativeGasUsed"`
	GasUsed           int64    `json:"gasUsed"`
	ContractAddress   string   `json:"contractAddress"`
	Logs              []string `json:"logs"`
}
