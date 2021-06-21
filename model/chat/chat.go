package chat


// ChatProtocolChatReqModel 文本聊天：请求
type ChatProtocolChatReqModel struct {
	Message		string	`json:"message" form:"message"`
	ToUID		int		`json:"to_uid" form:"to_uid"`
}

// ChatProtocolChatResModel 文本聊天：响应
type ChatProtocolChatResModel struct {
	Message		string	`json:"message" form:"message"`
	Datetime	string	`json:"datetime" form:"datetime"`
}
