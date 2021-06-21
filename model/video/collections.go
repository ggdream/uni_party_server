package video


// VideoUserCollectionsReqModel 获取用户收藏的视频：请求
type VideoUserCollectionsReqModel struct {
	UID    int `json:"uid" form:"uid"`
	Offset int `json:"offset" form:"offset"`
	Number int `json:"number" form:"number"`
}

// VideoUserCollectionsResModel 获取用户收藏的视频：响应
type VideoUserCollectionsResModel struct {
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
