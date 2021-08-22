package event

// SubscribeReqModel 获取订阅消息：请求
type SubscribeReqModel struct {
	Page int64 `json:"page" form:"page" binding:"required"`
}

// SubscribeResModel 获取订阅消息：响应
type SubscribeResModel struct {
	Unread int64         `json:"unread" form:"unread"`
	Result []EventResult `json:"result" form:"result"`
}
