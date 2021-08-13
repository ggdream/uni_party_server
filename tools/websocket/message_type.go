package websocket


const (
	// 文本消息
	TEXT ChatType = iota

	// 图片消息
	IMAGE
	// 音频消息
	AUDIO
	// 视频消息
	VIDEO

	// 全量同步
	FullSync
	// 增量同步
	IncrSync
)

// ChatType Websocket通信消息Enum类型
type ChatType int8

// NewChatType 将int8类型转为Enum类型
func NewChatType(index int8) ChatType {
	return ChatType(index)
}

// Index 将Enum类型转为int8类型
func (c ChatType) Index() int8 {
	return int8(c)
}
