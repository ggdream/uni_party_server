package helper


// ContactUsReqModel 联系我们：请求
type ContactUsReqModel struct {}

// ContactUsResModel 联系我们：响应
type ContactUsResModel struct {
	Phone string `json:"phone"`
	Email struct {
		Feedback       string `json:"feedback"`
		Authentication string `json:"authentication"`
		Cooperation    string `json:"cooperation"`
	} `json:"email"`
}
