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
 * @file eth-gettransactionreceipt_test.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */
package test

import (
	"strings"
	"testing"

	"github.com/fraymond/web3go"
	"github.com/fraymond/web3go/providers"
)

func TestEthGetTransactionReceipt(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:30305", 10, false))

	tx, err := connection.Eth.GetTransactionReceipt("0xfd53f3ccf3fb55b0333862b804abc6765d1433141b5a860e978f67794861a79b")

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if strings.Compare(tx.ContractAddress, "0x18a672e11d637fffadccc99b152f4895da069601") != 0 {
		t.Error("Invalid contract address")
		t.FailNow()
	}

}
