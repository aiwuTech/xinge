// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package xinge

import "errors"

type Client struct {
	AccessId  string
	AccessKey string
	ValidTime uint
	SecretKey string
}

func NewClient(accessId string, validTime uint, accessKey, secretKey string) *Client {
	return &Client{
		AccessId:  accessId,
		AccessKey: accessKey,
		ValidTime: validTime,
		SecretKey: secretKey,
	}
}

func (cli *Client) NewRequest(method, url string) *Request {
	return &Request{
		HttpMethod: method,
		HttpUrl:    url,
		Params:     make(map[string]interface{}),
		Client:     cli,
	}
}

func (cli *Client) AppDeviceNum() (int64, error) {
	request := cli.NewRequest("GET", deviceNumUrl)

	response, err := request.Execute()
	if err != nil {
		return 0, errors.New("<xinge> request app device num err:" + err.Error())
	}

	if !response.OK() {
		return 0, errors.New("<xinge> response err:" + response.Error())
	}

	return int64(response.Result["device_num"].(float64)), nil
}
