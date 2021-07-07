package event


// EventArchiveCollectReqModel 收藏：请求
type EventArchiveCollectReqModel struct {
	Eid  string `json:"eid" form:"eid"`
	Type int8   `json:"type" form:"type"`
}

// EventArchiveCollectResModel 收藏：响应
type EventArchiveCollectResModel struct {
	Status bool `json:"status" form:"status"`
}
