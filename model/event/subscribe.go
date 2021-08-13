package event


// SubscribeReqModel 获取订阅消息：请求
type SubscribeReqModel struct {
	Offset		int	`json:"offset" form:"offset"`
	Number		int	`json:"number" form:"number"`
}

// SubscribeResModel 获取订阅消息：响应
type SubscribeResModel struct {
	Unread		int	`json:"unread" form:"unread"`
	Events		[]eventResult	`json:"events" form:"events"`
}
