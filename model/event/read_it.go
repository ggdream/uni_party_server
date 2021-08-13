package event

// ArchiveGetReqModel 已读：请求
type ArchiveGetReqModel struct {
	Eid  string `json:"eid" form:"eid"`
	Type int8   `json:"type" form:"type"`
}

// ArchiveGetResModel 已读：响应
type ArchiveGetResModel struct {
	Status bool `json:"status" form:"status"`
}
