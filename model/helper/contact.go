package helper


// HelperContactUsReqModel 联系我们：请求
type HelperContactUsReqModel struct {}

// HelperContactUsResModel 联系我们：响应
type HelperContactUsResModel struct {
	Phone string `json:"phone"`
	Email struct {
		Feedback       string `json:"feedback"`
		Authentication string `json:"authentication"`
		Cooperation    string `json:"cooperation"`
	} `json:"email"`
}
