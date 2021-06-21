package archive


// CollectReqModel 收藏：请求
type CollectReqModel struct {
	VID			string		`json:"vid" form:"vid"`
	Type		int8		`json:"type" form:"type"`
}

// CollectResModel 收藏：响应
type CollectResModel struct {
	Status		string		`json:"status" form:"status"`
}
