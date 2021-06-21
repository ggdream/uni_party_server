package helper


// HelperFeedbackReqModel 用户反馈：请求
type HelperFeedbackReqModel struct {
	Type    string `json:"type" form:"type"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Email   string `json:"email" form:"email"`
}

// HelperFeedbackResModel 用户反馈：响应
type HelperFeedbackResModel struct {}
