package video


// VideoUploadReqModel 上传视频：请求
type VideoUploadReqModel struct {
	Title string   `json:"title" form:"title"`
	Tags  []string `json:"tags" form:"tags"`

	// File

	Location string `json:"location" form:"location"`
}

// VideoUploadResModel 上传视频：响应
type VideoUploadResModel struct {
	Vid		string	`json:"vid"`
}
