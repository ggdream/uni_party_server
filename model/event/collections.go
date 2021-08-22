package event

// GetCollectionReqModel 获取用户收藏的消息：请求
type GetCollectionReqModel struct {
	Offset int64 `json:"offset" form:"offset"`
	Number int64 `json:"number" form:"number"`
}

// GetCollectionResModel 获取用户收藏的消息：响应
type GetCollectionResModel struct {
	Total  int           `json:"total" form:"total"`
	Result []EventResult `json:"result" form:"result"`
}
