package user

// FollowersGetReqModel 获取用户的粉丝信息：请求
type FollowersGetReqModel struct {
	UID    string `json:"uid" form:"uid"`
	Offset int64  `json:"offset" form:"offset"`
	Number int64  `json:"number" form:"number"`
}

// FollowersGetResModel 获取用户的粉丝信息：响应
type FollowersGetResModel struct {
	Total int64      `json:"total" form:"total"`
	Users []UserInfo `json:"users" form:"users"`
}
