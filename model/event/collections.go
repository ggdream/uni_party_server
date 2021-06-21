package event


// EventUserCollectionsReqModel 获取用户收藏的消息：请求
type EventUserCollectionsReqModel struct {
	UID    int `json:"uid" form:"uid"`
	Offset int `json:"offset" form:"offset"`
	Number int `json:"number" form:"number"`
}

// EventUserCollectionsResModel 获取用户收藏的消息：响应
type EventUserCollectionsResModel struct {
	Total  int           `json:"total" form:"total"`
	Result []eventResult `json:"result" form:"result"`
}
