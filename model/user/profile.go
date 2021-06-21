package user


// UserInfoGetReqModel 获取用户信息：请求
type UserInfoGetReqModel struct {
	UID		int	`json:"uid" form:"uid"`
}

// UserInfoGetResModel 获取用户信息：响应
type UserInfoGetResModel struct {
	UID          int    `json:"uid"`
	Uname        string `json:"uname"`
	Birthday     string `json:"birthday"`
	Sex          string `json:"sex"`
	Avatar       string `json:"avatar"`
	Motto        string `json:"motto"`
	Level        int    `json:"level"`
	SubLevel     int    `json:"sub_level"`
	Type         int    `json:"type"`
	OrgName      string `json:"org_name"`
	Followers    int    `json:"followers"`
	Following    int    `json:"following"`
	EventCounter int    `json:"event_counter"`
	VideoCounter int    `json:"video_counter"`
	IsFollowing  bool   `json:"is_following"`
}


// UserInfoUpdateReqModel	更新个人信息：请求
type UserInfoUpdateReqModel struct {
	Uname    string `json:"uname" form:"uname"`
	Birthday string `json:"birthday" form:"birthday"`
	Sex      string `json:"sex" form:"sex"`
	Motto    string `json:"motto" form:"motto"`
}

// UserInfoUpdateResModel	更新个人信息：响应
type UserInfoUpdateResModel struct {}


// UserProfileAvatarUploadReqModel 用户上传头像：请求
type UserProfileAvatarUploadReqModel struct {
	// File	string	`json:"file" form:"file"`
}

// UserProfileAvatarUploadResModel 用户上传头像：响应
type UserProfileAvatarUploadResModel struct {
	URL string `json:"url"`
}
