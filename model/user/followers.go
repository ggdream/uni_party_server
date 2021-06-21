package user


// UserFollowersGetReqModel 获取用户的粉丝信息：请求
type UserFollowersGetReqModel struct {
	UID    string `json:"uid" form:"uid"`
	Offset string `json:"offset" form:"offset"`
	Number string `json:"number" form:"number"`
}

// UserFollowersGetResModel 获取用户的粉丝信息：响应
type UserFollowersGetResModel struct {
	Total int        `json:"total" form:"total"`
	Users []userInfo `json:"users" form:"users"`
}
