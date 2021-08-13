package event


// DeleteReqModel 删除消息：请求
type DeleteReqModel struct {
	EID		string	`json:"eid" form:"eid"`
}

// DeleteResModel 删除消息：响应
type DeleteResModel struct {}
