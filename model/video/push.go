package video

// PushReqModel 视频推送获取：请求
type PushReqModel struct {
	Number  string `json:"number"`
	Pointer string `json:"pointer"`
}

// PushResModel 视频推送获取：响应
type PushResModel struct {
	Next struct {
		Pointer string `json:"pointer"`
		HasNext bool   `json:"has_next"`
	} `json:"next"`
	Result []videoResult `json:"result"`
}
