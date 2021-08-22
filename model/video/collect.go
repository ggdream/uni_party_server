package video

// CollectReqModel 收藏视频：请求
type CollectReqModel struct {
	Vid string `json:"vid"`
}

// CollectResModel 收藏视频：响应
type CollectResModel struct{}

// CollectGetReqModel 获取用户收藏的视频：请求
type CollectGetReqModel struct {
	Page int64 `json:"page" form:"page"`
}

// CollectGetResModel 获取用户收藏的视频：响应
type CollectGetResModel struct {
	Total  int64              `json:"total"`
	Result []VideoResultModel `json:"result"`
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

// VideoResultModel 包含用户信息的视频结构
type VideoResultModel struct {
	VID         string `json:"vid"`
	UID         uint   `json:"uid"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Tags        string `json:"tags"`
	CollectTime int64  `json:"collect_time,omitempty"`
	CreateTime  int64  `json:"create_time"`
	Status      int8   `json:"status"`
}
