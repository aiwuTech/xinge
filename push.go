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
	"encoding/json"
	"errors"
	"time"
)

type ReqPush struct {
	PushType     PushType
	DeviceToken  string // for single-device push
	UserAccounts []string
	Tags         []string
	TagsOp       string
	MessageType  MessageType
	Message      interface{}
	ExpireTime   int
	SendTime     time.Time
	MultiPkgType MultiPkgType
	PushEnv      PushEnv
	PlatformType PlatformType
	LoopTimes    int
	LoopInterval int
	Cli          *Client
}

func (req *ReqPush) Push() error {
	var request *Request

	switch req.PushType {
	case PushType_single_device:
		request = req.Cli.NewRequest("GET", singleDeviceUrl)
		request.SetParam("device_token", req.DeviceToken)
	case PushType_single_account:
		request = req.Cli.NewRequest("GET", singleAccountUrl)
		request.SetParam("account", req.UserAccounts[0])
	case PushType_multi_account:
		request = req.Cli.NewRequest("GET", multiAccountUrl)

		accounts, err := json.Marshal(req.UserAccounts)
		if err != nil {
			return errors.New("<xinge> marshal account list err:" + err.Error())
		}
		request.SetParam("account_list", string(accounts))
	case PushType_all_device:
		request = req.Cli.NewRequest("GET", allDeviceUrl)
		//		request.SetParam("loop_times", req.LoopTimes)
		//		request.SetParam("loop_interval", req.LoopInterval)
	case PushType_tags_device:
		request = req.Cli.NewRequest("GET", tagsDeviceUrl)

		tags, err := json.Marshal(req.Tags)
		if err != nil {
			return errors.New("<xinge> marshal tag list err:" + err.Error())
		}
		request.SetParam("tags_list", string(tags))
		request.SetParam("tags_op", req.TagsOp)
		//		request.SetParam("loop_times", req.LoopTimes)
		//		request.SetParam("loop_interval", req.LoopInterval)
	default:
		return errors.New("<xinge> invalid request push type.")
	}

	request.SetParam("message_type", req.MessageType)

	message := ""
	switch req.PlatformType {
	case Platform_android:
		// message
		if androidMsg, ok := req.Message.(*AndroidMessage); ok {
			androidMessage, err := json.Marshal(androidMsg)
			if err != nil {
				return errors.New("<xinge> marshal android message err:" + err.Error())
			}

			message = string(androidMessage)
		} else {
			return errors.New("<xinge> invalid android message content.")
		}

	case Platform_ios:
		// message
		if iosMsg, ok := req.Message.(*IosMessage); ok {
			iosMessage, err := json.Marshal(iosMsg)
			if err != nil {
				return errors.New("<xinge> marshal ios message err:" + err.Error())
			}

			message = string(iosMessage)
		} else {
			return errors.New("<xinge> invalid ios message content.")
		}
	}
	request.SetParam("message", message)

	request.SetParam("expire_time", req.ExpireTime)
	request.SetParam("send_time", req.SendTime.Format("2006-01-03 15:33:34"))
	// multi_pkg
	request.SetParam("multi_pkg", req.MultiPkgType)
	// environment
	request.SetParam("environment", req.PushEnv)

	response, err := request.Execute()
	if err != nil {
		return errors.New("<xinge> request err:" + err.Error())
	}

	if !response.OK() {
		return errors.New("<xinge> response err:" + response.Error())
	}

	return nil
}
