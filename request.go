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

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/aiwuTech/httpclient"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Request struct {
	HttpMethod string
	HttpUrl    string
	Params     map[string]interface{}
	Client     *Client
}

func (req *Request) SetParam(name string, value interface{}) {
	req.Params[name] = value
}

func (req *Request) SetParams(params map[string]interface{}) {
	for k, v := range params {
		req.SetParam(k, v)
	}
}

func (req *Request) Execute() (*Response, error) {
	body, err := req.doRequestAndGetBody()
	if err != nil {
		return nil, err
	}

	log.Println(string(body))
	rspMsg := new(Response)
	if err := json.Unmarshal(body, rspMsg); err != nil {
		return nil, err
	}

	return rspMsg, nil
}

func (req *Request) doRequestAndGetBody() ([]byte, error) {
	urls := req.HttpUrl + "?" + req.queryString()

	rsp, err := httpclient.ForwardHttp(req.HttpMethod, urls, nil)
	if err != nil {
		return nil, err
	}

	body := httpclient.GetForwardHttpBody(rsp.Body)
	return body, nil
}

func (req *Request) queryString() string {
	reqParams := req.makeRequestParams()
	sign := req.md5Signature(reqParams)

	values := url.Values{}
	values.Add("sign", sign)
	for k, v := range reqParams {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	return values.Encode()
}

/**
内容签名。生成规则：
A）提取请求方法method（GET或POST）；
B）提取请求url信息，包括Host字段的IP或域名和URI的path部分，注意不包括Host的端口和Path的querystring。请在请求中带上Host字段，否则将视为无效请求。
	比如openapi.xg.qq.com/v2/push/single_device或者10.198.18.239/v2/push/single_device;
C）将请求参数（不包括sign参数）格式化成K=V方式，注意：计算sign时所有参数不应进行urlencode；
D）将格式化后的参数以K的字典序升序排列，拼接在一起，注意字典序中大写字母在前；
E）拼接请求方法、url、排序后格式化的字符串以及应用的secret_key；
F）将E形成字符串计算MD5值，形成一个32位的十六进制（字母小写）字符串，即为本次请求sign（签名）的值；
Sign=MD5($http_method$url$k1=$v1$k2=$v2$secret_key); 该签名值基本可以保证请求是合法者发送且参数没有被修改，但无法保证不被偷窥。
例如： POST请求到接口http://openapi.xg.qq.com/v2/push/single_device，
	有四个参数，access_id=123，timestamp=1386691200，Param1=Value1，Param2=Value2，secret_key为abcde。
	则上述E步骤拼接出的字符串为
		POSTopenapi.xg.qq.com/v2/push/single_deviceParam1=Value1Param2=Value2access_id=123timestamp=1386691200abcde，
	注意字典序中大写在前。计算出该字符串的MD5为ccafecaef6be07493cfe75ebc43b7d53，以此作为sign参数的值
*/
func (req *Request) md5Signature(params map[string]interface{}) string {
	origin := req.joinRequestParams(params)
	if origin == "" {
		return ""
	}

	urls, err := url.ParseRequestURI(req.HttpUrl)
	if err != nil {
		return ""
	}

	origin = req.HttpMethod + urls.Host + urls.Path + origin + req.Client.SecretKey

	c := md5.New()
	c.Write([]byte(origin))
	return strings.ToLower(fmt.Sprintf("%X", c.Sum(nil)))
}

func (req *Request) joinRequestParams(params map[string]interface{}) string {
	if params == nil || len(params) == 0 {
		return ""
	}

	keys := make([]string, 0)
	origin := ""
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		origin += key + fmt.Sprintf("=%v", params[key])
	}

	return origin
}

func (req *Request) makeRequestParams() map[string]interface{} {
	ps := make(map[string]interface{})

	ps["access_id"] = req.Client.AccessId
	ps["timestamp"] = time.Now().Unix()
	ps["valid_time"] = req.Client.ValidTime

	for k, v := range req.Params {
		ps[k] = v
	}

	return ps
}
