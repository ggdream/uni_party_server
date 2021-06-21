package event


// EventNoticeDetailReqModel 获取通知消息详情：请求
type EventNoticeDetailReqModel struct {
	Type		int8	`json:"type" form:"type"`
	EID			string	`json:"eid" form:"eid"`
}

// EventNoticeDetailResModel 获取通知消息详情：响应
type EventNoticeDetailResModel struct {
	// 消息的基本元信息
	eventResult

	Content		string				`json:"content" form:"content"`
	EventDetail	noticeEventDetail	`json:"event_detail" form:"event_detail"`
}

// noticeEventDetail 通知消息携带的所有内容
type noticeEventDetail struct {
	Files		[]file	`json:"files" form:"files"`
}

// file 通知消息携带文件的元信息
type file struct {
	Name		string	`json:"name" form:"name"`
	Type		string	`json:"type" form:"type"`
	Location	string	`json:"location" form:"location"`
}


// EventNoticeCreateReqModel 发布通知消息：请求
type EventNoticeCreateReqModel struct {
	Title		string	`json:"title" form:"title"`
	Content		string	`json:"content" form:"content"`
	Tags		[]string`json:"tags" form:"tags"`
	Files		[]string`json:"files" form:"files"`
}

// EventNoticeCreateResModel 发布通知消息：响应
type EventNoticeCreateResModel struct {
	Type		int8	`json:"type" form:"type"`
	EID			string	`json:"eid" form:"eid"`
}


// EventNoticeUpdateReqModel 修改通知消息：请求
type EventNoticeUpdateReqModel struct {
	EID			string	`json:"eid" form:"eid"`
	Title		string	`json:"title" form:"title"`
	Content		string	`json:"content" form:"content"`
	Tags		[]string`json:"tags" form:"tags"`
	Files		[]string`json:"files" form:"files"`
}

// EventNoticeUpdateResModel 修改通知消息：响应
type EventNoticeUpdateResModel struct {
	Frequency	int		`json:"frequency" form:"frequency"`
	EID			string	`json:"eid" form:"eid"`
}
