package user

import "mime/multipart"

// ProfileGetReqModel 获取用户信息：请求
type ProfileGetReqModel struct {
	UID uint `json:"uid" form:"uid"`
}

// ProfileGetResModel 获取用户信息：响应
type ProfileGetResModel struct {
	UID          uint   `json:"uid"`
	Username     string `json:"username"`
	Birthday     int64  `json:"birthday"`
	Sex          string `json:"sex"`
	Avatar       string `json:"avatar"`
	Motto        string `json:"motto"`
	Rank         uint8  `json:"rank"`
	SubRank      uint8  `json:"sub_rank"`
	Type         uint8  `json:"type"`
	Org          string `json:"org"`
	Followers    int    `json:"followers"`
	Following    int    `json:"following"`
	VideoCounter int    `json:"video_counter"`
	IsFollowing  bool   `json:"is_following"`
}

// ProfileSetReqModel	更新个人信息：请求
type ProfileSetReqModel struct {
	Username string `json:"username" form:"username"`
	Birthday string `json:"birthday" form:"birthday"`
	Sex      string `json:"sex" form:"sex"`
	Motto    string `json:"motto" form:"motto"`
}

// ProfileSetResModel	更新个人信息：响应
type ProfileSetResModel struct {
	Username string `json:"username" form:"username"`
	Birthday string `json:"birthday" form:"birthday"`
	Sex      string `json:"sex" form:"sex"`
	Motto    string `json:"motto" form:"motto"`
}

// AvatarUploadReqModel 用户上传头像：请求
type AvatarUploadReqModel struct {
	File *multipart.FileHeader `json:"file" form:"file"`
}

// AvatarUploadResModel 用户上传头像：响应
type AvatarUploadResModel struct {
	URL string `json:"url"`
}
