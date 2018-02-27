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
 * @file http-provider.go
 * @authors:
 *   Raymond Fu <fraymond@gmail.com>
 * @date Feb 2018
 */

package providers

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"encoding/json"

	"github.com/fraymond/web3go/providers/util"
)

type HTTPProvider struct {
	address string
	timeout int32
	secure  bool
}

func NewHTTPProvider(address string, timeout int32, secure bool) *HTTPProvider {
	provider := new(HTTPProvider)
	provider.address = address
	provider.timeout = timeout
	return provider
}

func (provider HTTPProvider) SendRequest(v interface{}, method string, params interface{}) error {

	bodyString := util.JSONRPCObject{Version: "2.0", Method: method, Params: params, ID: rand.Intn(100)}

	prefix := "http://"
	if provider.secure {
		prefix = "https://"
	}

	body := strings.NewReader(bodyString.AsJsonString())
	req, err := http.NewRequest("POST", prefix+provider.address, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	var netClient = &http.Client{
		Timeout: time.Second * time.Duration(provider.timeout),
	}

	resp, err := netClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var bodyBytes []byte

	if resp.StatusCode == 200 {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	return json.Unmarshal(bodyBytes, v)

}

func (provider HTTPProvider) Close() error { return nil }
