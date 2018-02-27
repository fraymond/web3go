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
 * @file block.go
 * @authors:
 *   Jérôme Laurens <jeromelaurens@gmail.com>
 * @date Feb 2018
 */

package dto

import (
	// "encoding/json"
	// "fmt"
	// "strconv"

	"github.com/fraymond/web3go/complex/types"
)

type Block struct {
	Number     types.ComplexIntResponse `json:"number"`
	Hash       string                   `json:"hash"`
	ParentHash string                   `json:"parentHash"`
	Nonce      types.ComplexIntResponse `json:"nonce"`
	Timestamp  types.ComplexIntResponse `json:"timestamp"`
}

// func (block *Block) UnmarshalJSON(d []byte) error {
// 	fmt.Println("Inside block.UnmarshalJSON")
// 	type T2 Block // create new type with same structure as BLock but without its method set!
// 	x := struct {
// 		T2               // embed
// 		Timestamp string `json:"timestamp"`
// 	}{T2: T2(*block)} // don't forget this, if you do and 't' already has some fields set you would lose them

// 	if err := json.Unmarshal(d, &x); err != nil {
// 		return err
// 	}
// 	*block = Block(x.T2)
// 	var err error = nil
// 	block.Timestamp, err = strconv.ParseInt(x.Timestamp, 0, 64)
// 	fmt.Printf("Timestamp en string %v, en int : %v", x.Timestamp, block.Timestamp)
// 	return err
// }
