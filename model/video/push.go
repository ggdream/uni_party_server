package video

// PushReqModel 视频推送获取：请求
type PushReqModel struct {
	Number  string `json:"number"`
	Pointer string `json:"pointer"` // 数据库id加密后的值，代表起始
}

// PushResModel 视频推送获取：响应
type PushResModel struct {
	Pointer string             `json:"pointer"`
	HasNext bool               `json:"has_next"`
	Result  []VideoResultModel `json:"result"`
}
