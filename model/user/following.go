package user


// UserFollowingGetResModel 获取用户的关注信息：请求
type UserFollowingGetReqModel struct {
	UID    string `json:"uid" form:"uid"`
	Offset string `json:"offset" form:"offset"`
	Number string `json:"number" form:"number"`
}

// UserFollowingGetResModel 获取用户的关注信息：响应
type UserFollowingGetResModel struct {
	Total int        `json:"total" form:"total"`
	Users []userInfo `json:"users" form:"users"`
}


// UserFollowingActReqModel 关注某个用户：请求
type UserFollowingActReqModel struct {
	UID  string `json:"uid" form:"uid"`
	Type int8   `json:"type" form:"type"`
}

// UserFollowingActResModel 关注某个用户：响应
type UserFollowingActResModel struct {
	NeedCaptcha	bool	`json:"need_captcha" form:"need_captcha"`
}
