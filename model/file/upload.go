package file


// FileUploadReqModel 上传文件：请求
type FileUploadReqModel struct {
	Type int8   `json:"type" form:"type"`
	// File  string `json:"fid" form:"fid"`
}

// FileUploadResModel 上传文件：响应
type FileUploadResModel struct {
	NeedCaptcha bool `json:"need_captcha" form:"need_captcha"`
}
