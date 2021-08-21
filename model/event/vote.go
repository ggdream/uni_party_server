package event

// VoteDetailResModel 获取投票消息详情：响应
type VoteDetailResModel struct {
	// 消息的基本元信息
	eventResult

	Content string     `json:"content" form:"content"`
	Detail  voteDetail `json:"event_detail" form:"event_detail"`
}

type voteDetail struct {
	Multiple           bool     `json:"multiple" form:"multiple"`
	AllowedNumber      int      `json:"allowed_number" form:"allowed_number"`
	IsOver             bool     `json:"is_over" form:"is_over"`
	Deadline           string   `json:"deadline" form:"deadline"`
	Options            []string `json:"options" form:"options"`
	ParticipatedNumber int      `json:"participated_number" form:"participated_number"`
	IsParticipated     bool     `json:"is_participated" form:"is_participated"`
	SelectedNo         []int    `json:"selected_no" form:"selected_no"`
	Result             []int    `json:"result" form:"result"`
}

// VoteCreateReqModel 发布投票消息：请求
type VoteCreateReqModel struct {
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	Tags          []string `json:"tags"`
	AllowedNumber int      `json:"allowed_number"`
	Deadline      int64    `json:"deadline"`
	Options       []string `json:"options"`
}

// VoteCreateResModel 发布投票消息：响应
type VoteCreateResModel struct {
	Type int8   `json:"type"`
	EID  string `json:"eid"`
}

// VoteUpdateReqModel 修改投票消息：请求
type VoteUpdateReqModel struct {
	EID           string   `json:"eid" form:"eid"`
	Title         string   `json:"title" form:"title"`
	Content       string   `json:"content" form:"content"`
	Tags          []string `json:"tags" form:"tags"`
	AllowedNumber int      `json:"allowed_number" form:"allowed_number"`
	Deadline      int64    `json:"deadline" form:"deadline"`
	Options       []string `json:"options" form:"options"`
}

// VoteUpdateResModel 修改投票消息：响应
type VoteUpdateResModel struct {
	Frequency int    `json:"frequency"`
	EID       string `json:"eid"`
}

// VoteDoReqModel 执行投票：请求
type VoteDoReqModel struct {
	EID     string `json:"eid" form:"eid"`
	Answers []bool `json:"answers" form:"answers"`
}

// VoteDoResModel 执行投票：响应
type VoteDoResModel struct {
	Total   int      `json:"total"`
	Options []string `json:"options"`
	Answers []int64  `json:"answers"`
}
