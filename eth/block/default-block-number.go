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
 * @file default-block-number.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */

package block

import "github.com/fraymond/web3go/complex/types"

// NUMBER - An integer block number
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#the-default-block-parameter
func NUMBER(blocknumber types.ComplexIntParameter) string {
	return blocknumber.ToHex()
}

const (
	// EARLIEST - Earliest block
	EARLIEST string = "earliest"
	// LATEST - latest block
	LATEST string = "latest"
	// PENDING - Pending block
	PENDING string = "pending"
)
