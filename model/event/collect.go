package event


// ArchiveCollectReqModel 收藏：请求
type ArchiveCollectReqModel struct {
	Eid  string `json:"eid" form:"eid"`
	Type int8   `json:"type" form:"type"`
}

// ArchiveCollectResModel 收藏：响应
type ArchiveCollectResModel struct {
	Status bool `json:"status" form:"status"`
}
