package user


// UserSearchReqModel 搜索用户：请求
type UserSearchReqModel struct {
	Type  int8   `json:"type"`
	Query string `json:"query"`
}

// UserSearchResModel 搜索用户：响应
type UserSearchResModel struct {
	Total string     `json:"total"`
	Match []userInfo `json:"match"`
}
