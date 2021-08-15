package video

// SearchReqModel 搜索视频：请求
type SearchReqModel struct {
	Keyword string   `json:"keyword" binding:"required"`
	Tags    []string `json:"tags"`
}

// SearchResModel 搜索视频：响应
type SearchResModel struct {
	Result []videoResult `json:"result"`
	Count  int64         `json:"count"`
}
