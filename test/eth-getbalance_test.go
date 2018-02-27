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
 * @file eth-getbalance_test.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */

package test

import (
	"testing"

	web3 "github.com/fraymond/web3go"
	"github.com/fraymond/web3go/eth/block"
	"github.com/fraymond/web3go/providers"
)

func TestEthGetBalance(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

	_, err := connection.Eth.ListAccounts()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	bal, err := connection.Eth.GetBalance("0xcEB0030d28C591Be1679bAe40CcD3fe7fBbBCe07", block.LATEST)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(bal)

}
