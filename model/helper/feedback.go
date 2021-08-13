package helper

// FeedbackReqModel 用户反馈：请求
type FeedbackReqModel struct {
	Type    int8   `json:"type" form:"type"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Email   string `json:"email" form:"email"`
}

// FeedbackResModel 用户反馈：响应
type FeedbackResModel struct{}
