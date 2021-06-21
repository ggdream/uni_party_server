package chat


// ChatProtocolNotifyReqModel 后端通知：请求
type ChatProtocolNotifyReqModel struct {
	Message		string	`json:"message" form:"message"`
}
