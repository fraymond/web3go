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
 * @file web3.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */

package web3

import (
	"github.com/fraymond/web3go/complex/types"
	"github.com/fraymond/web3go/dto"
	"github.com/fraymond/web3go/eth"
	"github.com/fraymond/web3go/net"
	"github.com/fraymond/web3go/personal"
	"github.com/fraymond/web3go/providers"
)

// Coin - Ethereum value unity value
const (
	Coin float64 = 1000000000000000000
)

// Web3 - The Web3 Module
type Web3 struct {
	Provider providers.ProviderInterface
	Eth      *eth.Eth
	Net      *net.Net
	Personal *personal.Personal
}

// NewWeb3 - Web3 Module constructor to set the default provider, Eth, Net and Personal
func NewWeb3(provider providers.ProviderInterface) *Web3 {
	web3 := new(Web3)
	web3.Eth = eth.NewEth(provider)
	web3.Net = net.NewNet(provider)
	web3.Personal = personal.NewPersonal(provider)
	web3.Provider = provider
	return web3
}

// ClientVersion - Returns the current client version.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#web3_clientversion
// Parameters:
//    - none
// Returns:
// 	  - String - The current client version
func (web Web3) ClientVersion() (string, error) {

	pointer := &dto.RequestResult{}

	err := web.Provider.SendRequest(pointer, "web3_clientVersion", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// Sha3 - Returns Keccak-256 (not the standardized SHA3-256) of the given data.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#web3_sha3
//    - DATA - the data to convert into a SHA3 hash
// Returns:
// 	  - DATA - The SHA3 result of the given string.
func (web Web3) Sha3(data types.ComplexString) (string, error) {

	params := make([]string, 1)
	params[0] = data.ToHex()

	pointer := &dto.RequestResult{}

	err := web.Provider.SendRequest(pointer, "web3_sha3", params)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}
