package event


// EventParticipationDetailReqModel 获取报名消息详情：请求
type EventParticipationDetailReqModel struct {
	Type		int8	`json:"type" form:"type"`
	EID			string	`json:"eid" form:"eid"`
}

// EventParticipationDetailResModel 获取报名消息详情：响应
type EventParticipationDetailResModel struct {
	// 消息的基本元信息
	eventResult

	Content		string				`json:"content" form:"content"`
	EventDetail	partEventDetail		`json:"event_detail" form:"event_detail"`
}

type partEventDetail struct {
	RequiredNumber		int			`json:"required_number" form:"required_number"`
	AllowedCancel		bool		`json:"allowed_cancel" form:"allowed_cancel"`
	IsOver				bool		`json:"is_over" form:"is_over"`
	Deadline			string		`json:"deadline" form:"deadline"`
	IsParticipated		bool		`json:"is_participated" form:"is_participated"`
	ParticipatedNumber	int			`json:"participated_number" form:"participated_number"`
	Result				[]int		`json:"result" form:"result"`
}


// EventParticipationCreateReqModel 发布报名消息：请求
type EventParticipationCreateReqModel struct {
	Title          string   `json:"title" form:"title"`
	Content        string   `json:"content" form:"content"`
	Tags           []string `json:"tags" form:"tags"`
	RequiredNumber int      `json:"required_number" form:"required_number"`
	AllowedCancel  bool     `json:"allowed_cancel" form:"allowed_cancel"`
	Deadline       string   `json:"deadline" form:"deadline"`
}

// EventParticipationCreateResModel 发布报名消息：响应
type EventParticipationCreateResModel struct {
	Type int8   `json:"type" form:"type"`
	EID  string `json:"eid" form:"eid"`
}


// EventParticipationUpdateReqModel 修改报名消息：请求
type EventParticipationUpdateReqModel struct {
	EID            string   `json:"eid" form:"eid"`
	Title          string   `json:"title" form:"title"`
	Content        string   `json:"content" form:"content"`
	Tags           []string `json:"tags" form:"tags"`
	RequiredNumber int      `json:"required_number" form:"required_number"`
	AllowedCancel  bool     `json:"allowed_cancel" form:"allowed_cancel"`
	Deadline       string   `json:"deadline" form:"deadline"`
}

// EventParticipationUpdateResModel 修改报名消息：响应
type EventParticipationUpdateResModel struct {
	Frequency int    `json:"frequency" form:"frequency"`
	EID       string `json:"eid" form:"eid"`
}
