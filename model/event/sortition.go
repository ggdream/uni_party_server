package event

// SortitionDetailResModel 获取随机消息详情：响应
type SortitionDetailResModel struct {
	// 消息的基本元信息
	EventResult

	Content string     `json:"content" form:"content"`
	Detail  sortDetail `json:"event_detail" form:"event_detail"`
}

type sortDetail struct {
	RequiredNumber     int    `json:"required_number" form:"required_number"`
	AllowedCancel      bool   `json:"allowed_cancel" form:"allowed_cancel"`
	IsOver             bool   `json:"is_over" form:"is_over"`
	Deadline           string `json:"deadline" form:"deadline"`
	IsParticipated     bool   `json:"is_participated" form:"is_participated"`
	IsSelected         bool   `json:"is_selected" form:"is_selected"`
	ParticipatedNumber int    `json:"participated_number" form:"participated_number"`
	Result             []int  `json:"result" form:"result"`
}

// SortitionCreateReqModel 发布随机消息：请求
type SortitionCreateReqModel struct {
	Title          string `json:"title" form:"title"`
	Content        string `json:"content" form:"content"`
	RequiredNumber int    `json:"required_number" form:"required_number"`
	AllowedCancel  bool   `json:"allowed_cancel" form:"allowed_cancel"`
	Deadline       int64  `json:"deadline" form:"deadline"`
}

// SortitionCreateResModel 发布随机消息：响应
type SortitionCreateResModel struct {
	Type int8   `json:"type" form:"type"`
	EID  string `json:"eid" form:"eid"`
}

// SortitionUpdateReqModel 修改随机消息：请求
type SortitionUpdateReqModel struct {
	EID            string   `json:"eid" form:"eid"`
	Title          string   `json:"title" form:"title"`
	Content        string   `json:"content" form:"content"`
	Tags           []string `json:"tags" form:"tags"`
	RequiredNumber int      `json:"required_number" form:"required_number"`
	AllowedCancel  bool     `json:"allowed_cancel" form:"allowed_cancel"`
	Deadline       int64    `json:"deadline" form:"deadline"`
}

// SortitionUpdateResModel 修改随机消息：响应
type SortitionUpdateResModel struct {
	Frequency int    `json:"frequency" form:"frequency"`
	EID       string `json:"eid" form:"eid"`
}

// SortitionJoinReqModel 申请参加随机：请求
type SortitionJoinReqModel struct {
	EID  string `json:"eid" form:"eid"`
	Type int    `json:"type" form:"type"`
}

// SortitionJoinResModel 申请参加随机：响应
type SortitionJoinResModel struct{}
