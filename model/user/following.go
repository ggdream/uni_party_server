package user

// FollowingGetReqModel 获取用户的关注信息：请求
type FollowingGetReqModel struct {
	UID    string `json:"uid" form:"uid"`
	Offset int64  `json:"offset" form:"offset"`
	Number int64  `json:"number" form:"number"`
}

// FollowingGetResModel 获取用户的关注信息：响应
type FollowingGetResModel struct {
	Total int64      `json:"total" form:"total"`
	Users []UserInfo `json:"users" form:"users"`
}

// FollowingActReqModel 关注某个用户：请求
type FollowingActReqModel struct {
	UID  uint `json:"uid" form:"uid"`
	Type int8 `json:"type" form:"type"`
}

// FollowingActResModel 关注某个用户：响应
type FollowingActResModel struct{}
