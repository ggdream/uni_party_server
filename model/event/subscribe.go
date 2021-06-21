package event


// EventSubReqModel 获取订阅消息：请求
type EventSubReqModel struct {
	Offset		int	`json:"offset" form:"offset"`
	Number		int	`json:"number" form:"number"`
}

// EventSubResModel 获取订阅消息：响应
type EventSubResModel struct {
	Unread		int	`json:"unread" form:"unread"`
	Events		[]eventResult	`json:"events" form:"events"`
}

