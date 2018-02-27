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
 * @file web3-getstorageat_test.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */

package test

import (
	"testing"

	web3 "github.com/fraymond/web3go"
	"github.com/fraymond/web3go/complex/types"
	"github.com/fraymond/web3go/eth/block"
	"github.com/fraymond/web3go/providers"
)

func TestEthGetStorageAt(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:30305", 10, false))

	accounts, err := connection.Eth.ListAccounts()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	position := types.ComplexIntParameter(1)

	storage, err := connection.Eth.GetStorageAt(accounts[0], position, block.LATEST)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(storage)

}
