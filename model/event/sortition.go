package event


// EventSortitionDetailReqModel 获取随机消息详情：请求
type EventSortitionDetailReqModel struct {
	Type int8   `json:"type" form:"type"`
	EID  string `json:"eid" form:"eid"`
}

// EventSortitionDetailResModel 获取随机消息详情：响应
type EventSortitionDetailResModel struct {
	// 消息的基本元信息
	eventResult

	Content     string          `json:"content" form:"content"`
	EventDetail sortEventDetail `json:"event_detail" form:"event_detail"`
}

type sortEventDetail struct {
	RequiredNumber     int    `json:"required_number" form:"required_number"`
	AllowedCancel      bool   `json:"allowed_cancel" form:"allowed_cancel"`
	IsOver             bool   `json:"is_over" form:"is_over"`
	Deadline           string `json:"deadline" form:"deadline"`
	IsParticipated     bool   `json:"is_participated" form:"is_participated"`
	IsSelected         bool   `json:"is_selected" form:"is_selected"`
	ParticipatedNumber int    `json:"participated_number" form:"participated_number"`
	Result             []int  `json:"result" form:"result"`
}


// EventSortitionCreateReqModel 发布随机消息：请求
type EventSortitionCreateReqModel struct {
	Title          string   `json:"title" form:"title"`
	Content        string   `json:"content" form:"content"`
	Tags           []string `json:"tags" form:"tags"`
	RequiredNumber int      `json:"required_number" form:"required_number"`
	AllowedCancel  bool     `json:"allowed_cancel" form:"allowed_cancel"`
	Deadline       string   `json:"deadline" form:"deadline"`
}

// EventSortitionCreateResModel 发布随机消息：响应
type EventSortitionCreateResModel struct {
	Type int8   `json:"type" form:"type"`
	EID  string `json:"eid" form:"eid"`
}


// EventSortitionUpdateReqModel 修改随机消息：请求
type EventSortitionUpdateReqModel struct {
	EID            string   `json:"eid" form:"eid"`
	Title          string   `json:"title" form:"title"`
	Content        string   `json:"content" form:"content"`
	Tags           []string `json:"tags" form:"tags"`
	RequiredNumber int      `json:"required_number" form:"required_number"`
	AllowedCancel  bool     `json:"allowed_cancel" form:"allowed_cancel"`
	Deadline       string   `json:"deadline" form:"deadline"`
}

// EventSortitionUpdateResModel 修改随机消息：响应
type EventSortitionUpdateResModel struct {
	Frequency int    `json:"frequency" form:"frequency"`
	EID       string `json:"eid" form:"eid"`
}
