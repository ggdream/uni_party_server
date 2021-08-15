package video

import "mime/multipart"

// UploadReqModel 上传视频：请求
type UploadReqModel struct {
	Title string                `json:"title" form:"title" binding:"required"`
	Tags  []string              `json:"tags" form:"tags" binding:"required"`
	Cover *multipart.FileHeader `json:"cover" form:"cover" binding:"required"`
	Video *multipart.FileHeader `json:"video" form:"video" binding:"required"`

	Location string `json:"location" form:"location"`
}

// UploadResModel 上传视频：响应
type UploadResModel struct {
	Vid string `json:"vid"`
}

// UploadGetReqModel 获取用户发布的视频：请求
type UploadGetReqModel struct {
	UID    uint `json:"uid" form:"uid"`
	Offset int  `json:"offset" form:"offset"`
	Number int  `json:"number" form:"number"`
}

// UploadGetResModel 获取用户发布的视频：响应
type UploadGetResModel struct {
	Total  int64             `json:"total"`
	Result []VideoResultLess `json:"result"`
}

// VideoResultLess 不包含用户信息的视频结构
type VideoResultLess struct {
	Vid            string   `json:"vid"`
	Title          string   `json:"title"`
	Cover          string   `json:"cover"`
	Tags           []string `json:"tags"`
	CreateTime     int64    `json:"create_time"`
	WatchCounter   string   `json:"watch_counter"`
	StarCounter    string   `json:"star_counter"`
	CommentCounter string   `json:"comment_counter"`
	IsGet          string   `json:"is_get"`
	IsCollect      string   `json:"is_collect"`
}
