package archive


// StarReqModel 点赞：请求
type StarReqModel struct {
	VID		string		`json:"vid" form:"vid"`
	Type	int8		`json:"type" form:"type"`
}

// StarResModel 点赞：响应
type StarResModel struct {
	Status	bool		`json:"status" form:"status"`
}
