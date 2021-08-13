package auth

// LoginCipherReqModel 密码登录：请求
type LoginCipherReqModel struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// LoginCipherResModel  密码登录：响应
type LoginCipherResModel struct {
	//NeedCaptcha bool   `json:"need_captcha" form:"need_captcha"`
	Token       string `json:"token" form:"token"`
}

// LoginCodeReqModel  验证码登录：请求
type LoginCodeReqModel struct {
	Account    string `json:"account" form:"account" binding:"required"`
	DeviceCode string `json:"device_code" form:"device_code" binding:"required"`
	SmsCode    string `json:"sms_code" form:"sms_code" binding:"required"`
}

// LoginCodeResModel  验证码登录：响应
type LoginCodeResModel struct {
	//NeedCaptcha bool   `json:"need_captcha" form:"need_captcha"`
	Token       string `json:"token" form:"token"`
}

// LoginSendPhoneCodeReqModel  发送手机验证码：请求
type LoginSendPhoneCodeReqModel struct {
	Phone string `json:"phone" form:"phone" binding:"required"`
}

// LoginSendPhoneCodeResModel  发送手机验证码：响应
type LoginSendPhoneCodeResModel struct {
	DeviceCode string `json:"device_code" form:"device_code"`
}

// LoginSendEMailCodeReqModel  发送邮箱验证码：请求
type LoginSendEMailCodeReqModel struct {
	EMail string `json:"email" form:"email" binding:"required"`
}

// LoginSendEMailCodeResModel  发送邮箱验证码：响应
type LoginSendEMailCodeResModel struct {
	DeviceCode string `json:"device_code" form:"device_code"`
}
