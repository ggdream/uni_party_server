package video


// VideoPushReqModel 视频推送获取：请求
type VideoPushReqModel struct {
	Number  string `json:"number"`
	Pointer string `json:"pointer"`
}

// VideoPushResModel 视频推送获取：响应
type VideoPushResModel struct {
	Pointer string        `json:"pointer"`
	Captcha bool          `json:"captcha"`
	Result  []videoResult `json:"result"`
}
