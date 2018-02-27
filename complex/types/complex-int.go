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
 * @file complex-int.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */

package types

import (
	"fmt"
	"strconv"
	"strings"
)

type ComplexIntParameter int64

func (s ComplexIntParameter) ToHex() string {

	return fmt.Sprintf("0x%x", s)

}

type ComplexIntResponse string

func (s ComplexIntResponse) ToUInt64() uint64 {

	stringValue := string(s)

	cleaned := strings.Replace(stringValue, "0x", "", -1)
	sResult, _ := strconv.ParseUint(cleaned, 16, 64)

	return sResult

}

func (s ComplexIntResponse) ToInt64() int64 {

	stringValue := string(s)

	cleaned := strings.Replace(stringValue, "0x", "", -1)
	sResult, _ := strconv.ParseInt(cleaned, 16, 64)

	return sResult

}
