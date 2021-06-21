package video


// VideoSearchReqModel 搜索视频：请求
type VideoSearchReqModel struct {
	SortType int    `json:"sort_type"`
	Query    string `json:"query"`
	Offset   int    `json:"offset"`
	Number   int    `json:"number"`
}

// VideoSearchResModel 搜索视频：响应
type VideoSearchResModel struct {
	Captcha bool          `json:"captcha"`
	Result  []videoResult `json:"result"`
}
