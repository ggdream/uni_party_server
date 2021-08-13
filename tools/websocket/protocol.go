package websocket

import "encoding/json"

// Wrapper 通信请求体
type Wrapper struct {
	// Type 消息类型
	Type int8 `json:"type"`

	// FromUID 发送消息的用户的UID
	FromUID int `json:"from_uid"`
	// ToUID 接受消息的用户的UID
	ToUID int `json:"to_uid,omitempty"`
	// GroupID 组聊ID
	GroupID string `json:"group_id,omitempty"`

	// Message 文本消息
	Message string `json:"message,omitempty"`
	// URL 二进制文件地址
	URL string `json:"url,omitempty"`
	// Others 扩展字段，附加的其他数据
	Others interface{} `json:"others,omitempty"`

	// MID 消息ID，服务器收到消息时，无此字段；转发消息时，需要添加此字段
	MID string `json:"mid,omitempty"`
	// Datetime 消息接受时间，服务器收到消息时，无此字段；转发消息时，需要添加此字段
	Datetime int `json:"datetime,omitempty"`
	// Version 为方便后期协议修改，添加版本号字段
	Version string `json:"version"`
	// Signature 内容签名，服务器收到消息时，有此字段；转发消息时，需要添加此字段
	Signature string `json:"signature,omitempty"`
}

// NewWrapper 将JSON字符串反序列化为Wrapper
func NewWrapper(data []byte) (wrapper *Wrapper, err error) {
	err = json.Unmarshal(data, wrapper)
	return
}

// Modify 修改原有反序列化号的内容为要转发的消息
func (w *Wrapper) Modify(datetime int, mid string) {
	// 清空无用字段
	w.Signature = ""

	w.Datetime = datetime
	w.MID = mid
}

// ToJSON 将Wrapper序列化为JSON字符串
func (w *Wrapper) ToJSON() ([]byte, error) {
	return json.Marshal(*w)
}
