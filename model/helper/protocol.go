package helper


// HelperProtocolUserReqModel 获取用户协议：请求
type HelperProtocolUserReqModel struct {
	Version		string	`json:"version" form:"version"`
}

// HelperProtocolUserResModel 获取用户协议：响应
type HelperProtocolUserResModel struct {
	Content		string	`json:"content" form:"content"`
}


// HelperProtocolContentReqModel 获取服务协议：请求
type HelperProtocolContentReqModel struct {
	Version		string	`json:"version" form:"version"`
}

// HelperProtocolContentResModel 获取服务协议：响应
type HelperProtocolContentResModel struct {
	Content		string	`json:"content" form:"content"`
}

