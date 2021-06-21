package archive


// ReplyAddReqModel 添加评论：请求
type ReplyAddReqModel struct {
	VID		string		`json:"vid" form:"vid"`
	Root	string		`json:"root" form:"root"`
	Parent	string		`json:"parent" form:"parent"`
	Message	string		`json:"message" form:"message"`
}

// ReplyAddResModel 添加评论：响应
type ReplyAddResModel struct {
	RID			string		`json:"rid" form:"rid"`
	Message		string		`json:"message" form:"message"`
	Datetime	string		`json:"datetime" form:"datetime"`
	NeedCaptcha	string		`json:"need_captcha" form:"need_captcha"`
}


// ReplyDelReqModel 删除评论：请求
type ReplyDelReqModel struct {
	VID			string		`json:"vid" form:"vid"`
	RID			string		`json:"rid" form:"rid"`
}

// ReplyDelResModel 删除评论：响应
type ReplyDelResModel struct {
	NeedCaptcha	string		`json:"need_captcha" form:"need_captcha"`
}
