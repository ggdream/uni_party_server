package event

// EventArchiveStarReqModel 给消息点赞（get）：请求
type EventArchiveStarReqModel struct {
	Eid  string `json:"eid" form:"eid"`
	Type int8   `json:"type" form:"type"`
}

// EventArchiveStarResModel 给消息点赞（get）：响应
type EventArchiveStarResModel struct {
	Status bool `json:"status" form:"status"`
}
