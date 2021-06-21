package event


// EventArchiveReplyAddReqModel 添加评论：请求
type EventArchiveReplyAddReqModel struct {
	Vid     string `json:"vid" form:"vid"`
	Root    string `json:"root" form:"root"`
	Parent  string `json:"parent" form:"parent"`
	Message string `json:"message" form:"message"`
}

// EventArchiveReplyAddResModel 添加评论：响应
type EventArchiveReplyAddResModel struct {
	Rid         string `json:"rid" form:"rid"`
	Message     string `json:"message" form:"message"`
	Datetime    string `json:"datetime" form:"datetime"`
	NeedCaptcha bool   `json:"need_captcha" form:"need_captcha"`
}

// EventArchiveReplyDelReqModel 删除评论：请求
type EventArchiveReplyDelReqModel struct {
	Vid string `json:"vid" form:"vid"`
	Rid string `json:"rid" form:"rid"`
}

// EventArchiveReplyDelResModel 删除评论：响应
type EventArchiveReplyDelResModel struct {
	NeedCaptcha bool `json:"need_captcha" form:"need_captcha"`
}
