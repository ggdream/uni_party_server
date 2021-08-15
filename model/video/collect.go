package video

// CollectReqModel 收藏视频：请求
type CollectReqModel struct {
	Vid string `json:"vid"`
}

// CollectResModel 收藏视频：响应
type CollectResModel struct{}

// CollectGetReqModel 获取用户收藏的视频：请求
type CollectGetReqModel struct {
	Offset int64 `json:"offset" form:"offset"`
	Number int64 `json:"number" form:"number"`
}

// CollectGetResModel 获取用户收藏的视频：响应
type CollectGetResModel struct {
	Total  string        `json:"total"`
	Result []videoResult `json:"result"`
}

// userInfo 用户基本信息
type userInfo struct {
	UID     string `json:"uid"`
	Uname   string `json:"uname"`
	Avatar  string `json:"avatar"`
	Sex     string `json:"sex"`
	Motto   string `json:"motto"`
	Type    int8   `json:"type"`
	OrgName string `json:"org_name"`
}

// videoResult 包含用户信息的视频结构
type videoResult struct {
	Vid            string   `json:"vid"`
	UserInfo       userInfo `json:"userinfo"`
	Title          string   `json:"title"`
	Cover          string   `json:"cover"`
	Tags           string   `json:"tags"`
	WatchCounter   string   `json:"watch_counter"`
	StarCounter    string   `json:"star_counter"`
	CommentCounter string   `json:"comment_counter"`
	CollectTime    string   `json:"collect_time"`
	CreateTime     string   `json:"create_time"`
	UpdateTime     string   `json:"update_time"`
	IsGet          string   `json:"is_get"`
	IsCollect      string   `json:"is_collect"`
}
