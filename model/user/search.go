package user


// SearchReqModel 搜索用户：请求
type SearchReqModel struct {
	Type  int8   `json:"type"`
	Query string `json:"query"`
}

// SearchResModel 搜索用户：响应
type SearchResModel struct {
	Total string     `json:"total"`
	Match []userInfo `json:"match"`
}
