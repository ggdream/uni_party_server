package chat


// ChatProtocolFullSyncReqModel 全量同步：请求
type ChatProtocolFullSyncReqModel struct {}

// ChatProtocolFullSyncResModel 全量同步：响应
type ChatProtocolFullSyncResModel struct {
	Total		int				`json:"total" form:"total"`
	Result		[]syncResult	`json:"result" form:"result"`
}


// ChatProtocolIncrSyncReqModel 增量同步：请求
type ChatProtocolIncrSyncReqModel struct {
	UID			string			`json:"uid" form:"uid"`
	Pointer		string			`json:"pointer" form:"pointer"`
	Number		int				`json:"number" form:"number"`
}

// ChatProtocolIncrSyncResModel 增量同步：响应
type ChatProtocolIncrSyncResModel struct {
	Total		int				`json:"total" form:"total"`
	Result		[]syncResult	`json:"result" form:"result"`
}


type syncResult struct {
	FromUID		int		`json:"from_uid" form:"from_uid"`
	Datetime	string	`json:"datetime" form:"datetime"`
	DID			string	`json:"did" form:"did"`
	Message		string	`json:"message" form:"message"`
}
