package video


// VideoUserPublicationsReqModel 获取用户发布的视频：请求
type VideoUserPublicationsReqModel struct {
	UID    int `json:"uid" form:"uid"`
	Offset int `json:"offset" form:"offset"`
	Number int `json:"number" form:"number"`
}

// VideoUserPublicationsResModel 获取用户发布的视频：响应
type VideoUserPublicationsResModel struct {
	Total  string        `json:"total"`
	Result []videoResultLess `json:"result"`
}

// videoResultLess 不包含用户信息的视频结构
type videoResultLess struct {
	Vid            string   `json:"vid"`
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
