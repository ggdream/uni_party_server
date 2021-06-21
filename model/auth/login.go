package auth


// AuthLoginCipherReqModel 密码登录：请求
type AuthLoginCipherReqModel struct {
	Account		string		`json:"account" form:"account"`
	Password	string		`json:"password" form:"password"`
}

// AuthLoginCipherResModel 密码登录：响应
type AuthLoginCipherResModel struct {
	NeedCaptcha	bool		`json:"need_captcha" form:"need_captcha"`
	Token		string		`json:"token" form:"token"`
}


// AuthLoginCodeReqModel 验证码登录：请求
type AuthLoginCodeReqModel struct {
	DeviceCode		string		`json:"device_code" form:"device_code"`
	SmsCode			string		`json:"sms_code" form:"sms_code"`
}

// AuthLoginCodeResModel 验证码登录：响应
type AuthLoginCodeResModel struct {
	NeedCaptcha	bool		`json:"need_captcha" form:"need_captcha"`
	Token		string		`json:"token" form:"token"`
}


// AuthLoginSendPhoneCodeReqModel 发送手机验证码：请求
type AuthLoginSendPhoneCodeReqModel struct {
	Phone	string	`json:"phone" form:"phone"`
}

// AuthLoginSendPhoneCodeResModel 发送手机验证码：响应
type AuthLoginSendPhoneCodeResModel struct {
	DeviceCode		string		`json:"device_code" form:"device_code"`
}


// AuthLoginSendEMailCodeReqModel 发送邮箱验证码：请求
type AuthLoginSendEMailCodeReqModel struct {
	EMail	string	`json:"email" form:"email"`
}

// AuthLoginSendEMailCodeResModel 发送邮箱验证码：响应
type AuthLoginSendEMailCodeResModel struct {
	DeviceCode		string		`json:"device_code" form:"device_code"`
}
