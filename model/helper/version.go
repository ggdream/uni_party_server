package helper


// HelperVersionReqModel 获取最新版本信息：请求
type HelperVersionReqModel struct{}

// HelperVersionResModel 获取最新版本信息：响应
type HelperVersionResModel struct {
	Version  string `json:"version"`
	Content  string `json:"content"`
	Datetime string `json:"datetime"`
}


// HelperVersionsReqModel 获取所有版本信息：请求
type HelperVersionsReqModel struct{}

// HelperVersionsResModel 获取所有版本信息：响应
type HelperVersionsResModel struct {
	Version string `json:"version"`
	Result  []struct {
		Version  string `json:"version"`
		Content  string `json:"content"`
		Datetime string `json:"datetime"`
	} `json:"result"`
}
