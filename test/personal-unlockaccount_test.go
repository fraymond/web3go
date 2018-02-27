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
 * @file personal-unlockaccount_test.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */
package test

import (
	"errors"
	"testing"

	"github.com/fraymond/web3go"
	"github.com/fraymond/web3go/providers"
)

func TestPersonalUnlockAccount(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

	accounts, err := connection.Personal.ListAccounts()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	result, err := connection.Personal.UnlockAccount(accounts[0], "password", 100)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if !result {
		t.Error(errors.New("Can't unlock account"))
		t.FailNow()
	}

}
