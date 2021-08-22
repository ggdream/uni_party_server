package video

// SearchReqModel 搜索视频：请求
type SearchReqModel struct {
	Keyword string `json:"keyword" form:"keyword" binding:"required"`
	Page    int    `json:"page" form:"page" binding:"required"`
}

// SearchResModel 搜索视频：响应
type SearchResModel struct {
	Result []VideoResultModel `json:"result"`
	Total  int                `json:"count"`
}
