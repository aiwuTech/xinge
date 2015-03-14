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

import "fmt"

var (
	apiDomain  = "openapi.xg.qq.com"
	apiVersion = "v2"
)

var (
	singleDeviceUrl  = fmt.Sprintf("http://%s/%s/push/single_device", apiDomain, apiVersion)
	singleAccountUrl = fmt.Sprintf("http://%s/%s/push/single_account", apiDomain, apiVersion)
	multiAccountUrl  = fmt.Sprintf("http://%s/%s/push/account_list", apiDomain, apiVersion)
	allDeviceUrl     = fmt.Sprintf("http://%s/%s/push/all_device", apiDomain, apiVersion)
	tagsDeviceUrl    = fmt.Sprintf("http://%s/%s/push/tags_device", apiDomain, apiVersion)
	deviceNumUrl     = fmt.Sprintf("http://%s/%s/application/get_app_device_num", apiDomain, apiVersion)
	appTagsUrl       = fmt.Sprintf("http://%s/%s/tags/query_app_tags", apiDomain, apiVersion)
	batchSetTagsUrl  = fmt.Sprintf("http://%s/%s/tags/batch_set", apiDomain, apiVersion)
	batchDelTagsUrl  = fmt.Sprintf("http://%s/%s/tags/batch_del", apiDomain, apiVersion)
	tokenTagsUrl     = fmt.Sprintf("http://%s/%s/tags/query_token_tags", apiDomain, apiVersion)
	tagTokensUrl     = fmt.Sprintf("http://%s/%s/tags/query_tag_token_num", apiDomain, apiVersion)
)

type PlatformType byte

const (
	Platform_ios PlatformType = iota
	Platform_android
)

type MessageType byte

const (
	MessageType_ios MessageType = iota
	MessageType_notify
	MessageType_passthrough
)

type PushEnv byte

const (
	PushEnv_android PushEnv = iota
	PushEnv_prod
	PushEnv_dev
)

type MultiPkgType byte

const (
	MultiPkg_pkg MultiPkgType = iota
	MultiPkg_aid
	MultiPkg_ios
)

type PushType byte

const (
	PushType_single_device PushType = iota
	PushType_single_account
	PushType_multi_account
	PushType_all_device
	PushType_tags_device
)

const (
	TagsOp_AND = "AND"
	TagsOp_OR  = "OR"
)
