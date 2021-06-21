package event


// EventArchiveStarReqModel 给消息点赞（get）：请求
type EventArchiveStarReqModel struct {
	Vid  string `json:"vid" form:"vid"`
	Type int8   `json:"type" form:"type"`
}

// EventArchiveStarResModel 给消息点赞（get）：响应
type EventArchiveStarResModel struct {
	Status bool `json:"status" form:"status"`
}
