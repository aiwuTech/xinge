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

type AndroidMessage struct {
	Title         string                 `json:"title"`                 // 标题，必填
	Content       string                 `json:"content"`               // 内容，必填
	AcceptTime    []*AcceptTime          `json:"accept_time,omitempty"` //表示消息将在哪些时间段允许推送给用户，选填
	NotifyId      byte                   `json:"n_id,omitempty"`        //通知id，选填。若大于0，则会覆盖先前弹出的相同id通知；若为0，展示本条通知且不影响其他通知；若为-1，将清除先前弹出的所有通知，仅展示本条通知。默认为0
	BuilderId     int                    `json:"builder_id,omitempty"`  // 本地通知样式，必填
	Ring          byte                   `json:"ring,omitempty"`        // 是否响铃，0否，1是，下同。选填，默认1
	RingRaw       string                 `json:"ring_raw,omitempty"`    // 指定应用内的声音（ring.mp3），选填
	Vibrate       byte                   `json:"vibrate,omitempty"`     // 是否振动，选填，默认1
	Lights        byte                   `json:"lights,omitempty"`      // 是否呼吸灯，0否，1是，选填，默认1
	Clearable     byte                   `json:"clearable,omitempty"`   // 通知栏是否可清除，选填，默认1
	IconType      byte                   `json:"icon_type,omitempty"`   //默认0，通知栏图标是应用内图标还是上传图标,0是应用内图标，1是上传图标,选填
	IconRes       string                 `json:"icon_res,omitempty"`    // 应用内图标文件名（xg.png）或者下载图标的url地址，选填
	StyleId       byte                   `json:"style_id,omitempty"`    //Web端设置是否覆盖编号的通知样式，默认1，0否，1是,选填
	SmailIcon     string                 `json:"smail_icon,omitempty"`  //指定状态栏的小图片(xg.png),选填
	Action        *AndroidAction         `json:"action,omitempty"`      // 动作，选填。默认为打开app
	CustomContent map[string]interface{} `json:"custom_content"`
}

type AcceptTime struct {
	Start *HourMin `json:"start"`
	End   *HourMin `json:"end"`
}

type HourMin struct {
	Hour string `json:"hour"`
	Min  string `json:"min"`
}

type AndroidAction struct {
	ActionType  byte          `json:"action_type"` // 动作类型，1打开activity或app本身，2打开浏览器，3打开Intent，4通过包名拉起其他应用
	Activity    string        `json:"activity"`
	AtyAttr     *ActivityAttr `json:"aty_attr"` // activity属性，只针对action_type=1的情况
	Browser     *Browser      `json:"browser"`
	IntenterNet string        `json:"intent"`
	PackageName *Package      `json:"package_name"`
}

type ActivityAttr struct {
	IF byte `json:"if"` // 创建通知时，intent的属性，如：intent.setFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_RESET_TASK_IF_NEEDED);
	PF byte `json:"pf"` // PendingIntent的属性，如：PendingIntent.FLAG_UPDATE_CURRENT
}

// url：打开的url，confirm是否需要用户确认
type Browser struct {
	Url     string `json:"url"`
	Confirm byte   `json:"confirm"`
}

type Package struct {
	PackageName  string `json:"packageName"`        // 要拉起的别的应用的包名
	PackageDLUrl string `json:"packageDownloadUrl"` //拉起应用的下载链接（若客户端没有找到此应用会自动去下载）
	Confirm      byte   `json:"confirm"`            //是否确认
}

type IosMessage struct {
	Aps           *ApsAttr               `json:"aps"`
	CustomContent map[string]interface{} `json:"custom_content,omitempty"` //参考android的自定义属性
}

type ApsAttr struct {
	Alert string `json:"alert"`
	Badge string `json:"badge,omitempty"`
	Sound string `json:"sound,omitempty"`
}
