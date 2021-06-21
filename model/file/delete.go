package file


// FileDeleteReqModel 删除文件：请求
type FileDeleteReqModel struct {
	Type int8   `json:"type" form:"type"`
	FID  string `json:"fid" form:"fid"`
}

// FileDeleteReqModel 删除文件：响应
type FileDeleteResModel struct {
	NeedCaptcha bool `json:"need_captcha" form:"need_captcha"`
}
