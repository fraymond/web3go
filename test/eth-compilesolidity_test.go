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
 * @file eth-compilesolidity_test.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */
package test

import (
	"testing"

	"github.com/fraymond/web3go"
	"github.com/fraymond/web3go/providers"
)

func TestEthCompileSolidity(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

	code := "var greeterSource = 'contract mortal { address owner; function mortal() { owner = msg.sender; } function kill() { if (msg.sender == owner) selfdestruct(owner); } } contract greeter is mortal { string greeting; function greeter(string _greeting) public { greeting = _greeting; } function greet() constant returns (string) { return greeting; } }"

	compiled, err := connection.Eth.CompileSolidity(code)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(compiled.ToString())

}
