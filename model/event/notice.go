package event

import "mime/multipart"

// NoticeDetailReqModel 获取通知消息详情：请求
type NoticeDetailReqModel struct {
	Type int8   `json:"type" form:"type"`
	EID  string `json:"eid" form:"eid"`
}

// NoticeDetailResModel 获取通知消息详情：响应
type NoticeDetailResModel struct {
	// 消息的基本元信息
	eventResult

	Content string       `json:"content" form:"content"`
	Detail  noticeDetail `json:"event_detail" form:"event_detail"`
}

// noticeDetail 通知消息携带的所有内容
type noticeDetail struct {
	Files []file `json:"files" form:"files"`
}

// file 通知消息携带文件的元信息
type file struct {
	Name     string `json:"name" form:"name"`
	Type     string `json:"type" form:"type"`
	Location string `json:"location" form:"location"`
}

// NoticeCreateReqModel 发布通知消息：请求
type NoticeCreateReqModel struct {
	Title   string                  `json:"title" form:"title" binding:"required"`
	Content string                  `json:"content" form:"content" binding:"required"`
	Tags    []string                `json:"tags" form:"tags" binding:"required"`
	Files   []*multipart.FileHeader `json:"files" form:"files" binding:"required"`
}

// NoticeCreateResModel 发布通知消息：响应
type NoticeCreateResModel struct {
	Type int8   `json:"type" form:"type"`
	EID  string `json:"eid" form:"eid"`
}

// NoticeUpdateReqModel 修改通知消息：请求
type NoticeUpdateReqModel struct {
	EID     string   `json:"eid" form:"eid" binding:"required"`
	Title   string   `json:"title" form:"title" binding:"required"`
	Content string   `json:"content" form:"content" binding:"required"`
	Tags    []string `json:"tags" form:"tags" binding:"required"`
	//Files   []string `json:"files" form:"files" binding:"required"`
}

// NoticeUpdateResModel 修改通知消息：响应
type NoticeUpdateResModel struct {
	Frequency int    `json:"frequency" form:"frequency"`
	EID       string `json:"eid" form:"eid"`
}
