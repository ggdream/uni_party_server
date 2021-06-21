package helper


// HelperAboutUsReqModel 关于我们：请求
type HelperAboutUsReqModel struct {}

// HelperAboutUsResModel 关于我们：响应
type HelperAboutUsResModel struct {
	Content string `json:"content" form:"content"`
}
