package event

// SearchReqModel 搜索消息：请求
type SearchReqModel struct {
	//Type  int8   `json:"type" form:"type"`
	Query string `json:"query" form:"query"`

	Offset int `json:"offset" form:"offset"`
	Number int `json:"number" form:"number"`
}

// SearchResModel 搜索消息：响应
type SearchResModel struct {
	Total int           `json:"total"`
	Match []EventResult `json:"match"`
}

type EventResultModel struct {
	UID        uint   `json:"uid"`
	EID        string `json:"eid"`
	Title      string `json:"title" `
	Type       int8   `json:"type"`
	CreateTime int64  `json:"create_time"`
}
