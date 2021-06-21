package archive


// ArchiveBehaviorSearchReqModel 用户搜索行为数据搜集：请求
type ArchiveBehaviorSearchReqModel struct {
	Query string `json:"query" form:"form"`
}

// ArchiveBehaviorSearchResModel 用户搜索行为数据搜集：响应
type ArchiveBehaviorSearchResModel struct {}


// ArchiveBehaviorTimeReqModel 用户时长行为数据搜集：请求
type ArchiveBehaviorTimeReqModel struct {
	VID string `json:"vid" form:"vid"`
}

// ArchiveBehaviorTimeResModel 用户时长行为数据搜集：响应
type ArchiveBehaviorTimeResModel struct {}
