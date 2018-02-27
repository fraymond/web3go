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
 * @file eth-getBlockByNumber_test.go
 * @authors:
 *   Jérôme Laurens <jeromelaurens@gmail.com>
 * @date Feb 2018
 */

package test

import (
	"strings"
	"testing"
	"time"

	web3 "github.com/fraymond/web3go"
	"github.com/fraymond/web3go/complex/types"
	"github.com/fraymond/web3go/providers"
)

func TestEthGetBlockByNumber(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))

	//previous block Dec-09-2017 10:28:29 AM +UTC
	//Dec-09-2017 10:28:31 AM +UTC should be block 4701754, hash 0xaec14e98d578351a23d5ab40f65ee431063f582b5d736bbc0705a2eef0fb8f9d
	//next block Dec-09-2017 10:28:55 AM +UTC
	expectedBlockHash := "0xaec14e98d578351a23d5ab40f65ee431063f582b5d736bbc0705a2eef0fb8f9d"
	var blockNumber types.ComplexIntParameter = 4701754

	block, err := connection.Eth.GetBlockByNumber(blockNumber, false)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if block == nil {
		t.Error("Block returned is nil")
		t.FailNow()
	}

	actualBlockDate := time.Unix(block.Timestamp.ToInt64(), 0)
	expectedBlockDate := time.Date(2017, 12, 9, 10, 28, 31, 0, time.UTC)

	if strings.Compare(block.Hash, expectedBlockHash) != 0 {
		t.Errorf("Expected block hash %v, got %v", expectedBlockHash, block.Hash)
		t.FailNow()
	}
	if block.Number.ToInt64() != int64(blockNumber) {
		t.Errorf("Expected block number %v, got %v", blockNumber, block.Number)
		t.FailNow()
	}
	if actualBlockDate.Equal(expectedBlockDate) {
		t.Errorf("Expected block date %v, got %v", expectedBlockDate, actualBlockDate)
		t.Fail()
	}
}
