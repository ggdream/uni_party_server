package event


// EventDeleteReqModel 删除消息：请求
type EventDeleteReqModel struct {
	Type	int8	`json:"type" form:"type"`
	EID		string	`json:"eid" form:"eid"`
}

// EventDeleteResModel 删除消息：响应
type EventDeleteResModel struct {}
