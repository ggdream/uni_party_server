package event

// GetPublicationReqModel 获取用户发布的消息：请求
type GetPublicationReqModel struct {
	Offset int64 `json:"offset" form:"offset"`
	Number int64 `json:"number" form:"number"`
}

// GetPublicationResModel 获取用户发布的消息：响应
type GetPublicationResModel struct {
	Total  int             `json:"total" form:"total"`
	Result []publishResult `json:"result" form:"result"`
}

// publishResult 用户发布的消息元数据（相比eventResult{}，没有UserInfo字段）
type publishResult struct {
	EID            string   `json:"eid" form:"eid"`
	Title          string   `json:"title" form:"title"`
	Type           int8     `json:"type" form:"type"`
	Tags           []string `json:"tags" form:"tags"`
	WatchCounter   int      `json:"watch_counter" form:"watch_counter"`
	GetCounter     int      `json:"get_counter" form:"get_counter"`
	CommentCounter int      `json:"comment_counter" form:"comment_counter"`
	CreateTime     string   `json:"create_time" form:"create_time"`
	UpdateTime     string   `json:"update_time" form:"update_time"`
	IsGet          bool     `json:"is_get" form:"is_get"`
	IsCollect      bool     `json:"is_collect" form:"is_collect"`
}
