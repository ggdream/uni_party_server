package auth


// AuthSafetyEnvReqModel 环境审核：请求
type AuthSafetyEnvReqModel struct {
	Signature	string	`json:"signature" form:"signature"`
}

// AuthSafetyEnvResModel 环境审核：响应
type AuthSafetyEnvResModel struct {}
