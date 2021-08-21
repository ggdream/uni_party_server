package event

// ParticipationDetailResModel 获取报名消息详情：响应
type ParticipationDetailResModel struct {
	// 消息的基本元信息
	eventResult

	Content string     `json:"content" form:"content"`
	Detail  partDetail `json:"event_detail" form:"event_detail"`
}

type partDetail struct {
	RequiredNumber     int    `json:"required_number" form:"required_number"`
	AllowedCancel      bool   `json:"allowed_cancel" form:"allowed_cancel"`
	IsOver             bool   `json:"is_over" form:"is_over"`
	Deadline           string `json:"deadline" form:"deadline"`
	IsParticipated     bool   `json:"is_participated" form:"is_participated"`
	ParticipatedNumber int    `json:"participated_number" form:"participated_number"`
	Result             []int  `json:"result" form:"result"`
}

// ParticipationCreateReqModel 发布报名消息：请求
type ParticipationCreateReqModel struct {
	Title          string   `json:"title" form:"title"`
	Content        string   `json:"content" form:"content"`
	Tags           []string `json:"tags" form:"tags"`
	RequiredNumber int      `json:"required_number" form:"required_number"`
	AllowedCancel  bool     `json:"allowed_cancel" form:"allowed_cancel"`
	Deadline       int64    `json:"deadline" form:"deadline"`
}

// ParticipationCreateResModel 发布报名消息：响应
type ParticipationCreateResModel struct {
	Type int8   `json:"type" form:"type"`
	EID  string `json:"eid" form:"eid"`
}

// ParticipationUpdateReqModel 修改报名消息：请求
type ParticipationUpdateReqModel struct {
	EID            string   `json:"eid" form:"eid"`
	Title          string   `json:"title" form:"title"`
	Content        string   `json:"content" form:"content"`
	Tags           []string `json:"tags" form:"tags"`
	RequiredNumber int      `json:"required_number" form:"required_number"`
	AllowedCancel  bool     `json:"allowed_cancel" form:"allowed_cancel"`
	Deadline       int64    `json:"deadline" form:"deadline"`
}

// ParticipationUpdateResModel 修改报名消息：响应
type ParticipationUpdateResModel struct {
	Frequency int    `json:"frequency" form:"frequency"`
	EID       string `json:"eid" form:"eid"`
}

// ParticipationJoinReqModel 申请参加报名：请求
type ParticipationJoinReqModel struct {
	EID  string `json:"eid" form:"eid"`
	Type int    `json:"type" form:"type"`
}

// ParticipationJoinResModel 申请参加报名：响应
type ParticipationJoinResModel struct{}
