package chat


// ChatProtocolPingReqModel 心跳监测：请求
type ChatProtocolPingReqModel struct {
	Message		string	`json:"message" form:"message"`
}

// ChatProtocolPingResModel 心跳监测：响应
type ChatProtocolPingResModel struct {
	Message		string	`json:"message" form:"message"`
}
