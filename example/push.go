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
package main

import (
	"fmt"
	"github.com/aiwuTech/xinge"
	"time"
)

var (
	xingeClient = xinge.NewClient("2100094636", 300, "A61ULFG44E8R", "3f4ab60b6c2d95d09587dd8fa17ac04b")
)

func main() {

	message := &xinge.AndroidMessage{
		Content: "wifi密码qq发过来",
		Title:   "wifi密码是多少",
	}

	reqPush := &xinge.ReqPush{
		PushType:     xinge.PushType_tags_device,
		Tags:         []string{"标签1", "标签2"},
		TagsOp:       xinge.TagsOp_AND,
		UserAccounts: []string{"sdfsadf", "lsdjfoiwjo"},
		DeviceToken:  "sdfsdfewfsadfsdfsdf",
		MessageType:  xinge.MessageType_notify,
		Message:      message,
		ExpireTime:   300,
		SendTime:     time.Now(),
		MultiPkgType: xinge.MultiPkg_aid,
		PushEnv:      xinge.PushEnv_android,
		PlatformType: xinge.Platform_android,
		LoopTimes:    2,
		LoopInterval: 7,
		Cli:          xingeClient,
	}
	fmt.Println(reqPush.Push())

	fmt.Println(xingeClient.AppDeviceNum())
}
