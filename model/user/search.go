package user

// SearchReqModel 搜索用户：请求
type SearchReqModel struct {
	//Type  int8   `json:"type"`
	Query string `json:"query" form:"query"`
	Page  int    `json:"page" form:"page"`
}

// SearchResModel 搜索用户：响应
type SearchResModel struct {
	Total int                   `json:"total"`
	Match []SimpleUserInfoModel `json:"match"`
}

type SimpleUserInfoModel struct {
	UID      uint   `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	College  string `json:"college"`
}
