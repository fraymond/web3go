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
 * @file complex-string.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */

package types

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type ComplexString string

func (s ComplexString) ToHex() string {

	return fmt.Sprintf("0x%x", s)

}

func (s ComplexString) ToString() string {

	stringValue := string(s)

	cleaned := strings.Replace(stringValue, "0x", "", -1)
	sResult, _ := hex.DecodeString(cleaned)

	return string(sResult)

}
