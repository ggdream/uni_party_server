package file


// UploadReqModel 上传文件：请求
type UploadReqModel struct {
	Type int8   `json:"type" form:"type"`
	//   string `json:"fid" form:"fid"`
}

// UploadResModel 上传文件：响应
type UploadResModel struct {
	NeedCaptcha bool `json:"need_captcha" form:"need_captcha"`
}
