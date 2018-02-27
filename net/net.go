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
 * @file net.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */

package net

import (
	"github.com/fraymond/web3go/complex/types"
	"github.com/fraymond/web3go/dto"
	"github.com/fraymond/web3go/providers"
)

// Net - The Net Module
type Net struct {
	provider providers.ProviderInterface
}

// NewNet - Net Module constructor to set the default provider
func NewNet(provider providers.ProviderInterface) *Net {
	net := new(Net)
	net.provider = provider
	return net
}

// IsListening - Returns true if client is actively listening for network connections.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#net_listening
// Parameters:
//    - none
// Returns:
// 	  - Boolean - true when listening, otherwise false.
func (net *Net) IsListening() (bool, error) {

	pointer := &dto.RequestResult{}

	err := net.provider.SendRequest(pointer, "net_listening", nil)

	if err != nil {
		return false, err
	}

	return pointer.ToBoolean()

}

// GetPeerCount - Returns number of peers currently connected to the client.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#net_peercount
// Parameters:
//    - none
// Returns:
// 	  - QUANTITY - integer of the number of connected peers.
func (net *Net) GetPeerCount() (types.ComplexIntResponse, error) {

	pointer := &dto.RequestResult{}

	err := net.provider.SendRequest(pointer, "net_peerCount", nil)

	if err != nil {
		return types.ComplexIntResponse(0), err
	}

	return pointer.ToComplexIntResponse()

}

// GetVersion - Returns the current network id.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#net_version
// Parameters:
//    - none
// Returns:
//    - String - The current network id.
//    "1": Ethereum Mainnet
//    "2": Morden Testnet (deprecated)
//    "3": Ropsten Testnet
//    "4": Rinkeby Testnet
//    "42": Kovan Testnet
func (net *Net) GetVersion() (string, error) {

	pointer := &dto.RequestResult{}

	err := net.provider.SendRequest(pointer, "net_version", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}
