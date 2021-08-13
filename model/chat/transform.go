package chat

// Transform 客户端请求连接服务器，请求协议转换
type Transform struct {
	Token      string `json:"token" form:"token" binding:"required"`
	DeviceType string `json:"device_type" form:"device_type"`
	Signature  string `json:"signature" form:"signature"`
}
