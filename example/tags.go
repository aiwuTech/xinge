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
	"github.com/globalways/xinge"
)

var (
	xingeClient = xinge.NewClient("2100094636", 300, "A61ULFG44E8R", "3f4ab60b6c2d95d09587dd8fa17ac04b")
)

func main() {
	fmt.Println(xingeClient.AppTags(0, 100))
	fmt.Println(xingeClient.SetTags([2]string{"标签1", "094858ba12dd24cf4856bff6f7696826c546bcda"}, [2]string{"标签2", "094858ba12dd24cf4856bff6f7696826c546bcda"}))
	fmt.Println(xingeClient.TokenTags("094858ba12dd24cf4856bff6f7696826c546bcda"))
	fmt.Println(xingeClient.TagTokensNum("标签1"))
	fmt.Println(xingeClient.DelTags([2]string{"标签1", "094858ba12dd24cf4856bff6f7696826c546bcda"}, [2]string{"标签2", "094858ba12dd24cf4856bff6f7696826c546bcda"}))
}
