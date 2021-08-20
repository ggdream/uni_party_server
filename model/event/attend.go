package event

// ArchiveAddAttendReqModel 添加关注：请求
type ArchiveAddAttendReqModel struct {
	Eid string `json:"eid" form:"eid"`
}

// ArchiveAddAttendResModel 添加关注：响应
type ArchiveAddAttendResModel struct{}

// ArchiveDelAttendReqModel 取消关注：请求
type ArchiveDelAttendReqModel struct {
	Eid  string `json:"eid" form:"eid"`
	Type int8   `json:"type" form:"type"`
}

// ArchiveDelAttendResModel 取消关注：响应
type ArchiveDelAttendResModel struct{}
