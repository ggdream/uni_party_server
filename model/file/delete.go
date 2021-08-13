package file


// DeleteReqModel 删除文件：请求
type DeleteReqModel struct {
	Type int8   `json:"type" form:"type"`
	FID  string `json:"fid" form:"fid"`
}

// DeleteResModel 删除文件：响应
type DeleteResModel struct {
	NeedCaptcha bool `json:"need_captcha" form:"need_captcha"`
}
