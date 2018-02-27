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
 * @file ipc-provider_test.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */
package test

import (
	"testing"

	web3 "github.com/fraymond/web3go"
	"github.com/fraymond/web3go/providers"
)

func Test_IPCProvider(t *testing.T) {

	var ethClient = web3.NewWeb3(providers.NewIPCProvider("/tmp/ethereum_dev_mode/geth.ipc"))

	var _, error = ethClient.ClientVersion()

	if error != nil {
		t.Error(error)
		t.Fail()
	}

}
