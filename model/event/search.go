package event


// SearchReqModel 搜索消息：请求
type SearchReqModel struct {
	Type      int8   `json:"type" form:"type"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
	Query     string `json:"query" form:"query"`
}

// SearchResModel 搜索消息：响应
type SearchResModel struct {
	Total string        `json:"total" form:"total"`
	Match []eventResult `json:"match" form:"match"`
}
