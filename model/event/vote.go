package event


// EventVoteDetailReqModel 获取投票消息详情：请求
type EventVoteDetailReqModel struct {
	Type		int8	`json:"type" form:"type"`
	EID			string	`json:"eid" form:"eid"`
}

// EventVoteDetailResModel 获取投票消息详情：响应
type EventVoteDetailResModel struct {
	// 消息的基本元信息
	eventResult

	Content		string				`json:"content" form:"content"`
	EventDetail	voteEventDetail		`json:"event_detail" form:"event_detail"`
}

type voteEventDetail struct {
	Multiple			bool		`json:"multiple" form:"multiple"`
	AllowedNumber		int			`json:"allowed_number" form:"allowed_number"`
	IsOver				bool		`json:"is_over" form:"is_over"`
	Deadline			string		`json:"deadline" form:"deadline"`
	Options				[]string	`json:"options" form:"options"`
	ParticipatedNumber	int			`json:"participated_number" form:"participated_number"`
	IsParticipated		bool		`json:"is_participated" form:"is_participated"`
	SelectedNo			[]int		`json:"selected_no" form:"selected_no"`
	Result				[]int		`json:"result" form:"result"`
}


// EventVoteCreateReqModel 发布投票消息：请求
type EventVoteCreateReqModel struct {
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	Tags          []string `json:"tags"`
	AllowedNumber string   `json:"allowed_number"`
	Deadline      string   `json:"deadline"`
	Options       []string `json:"options"`
}

// EventVoteCreateResModel 发布投票消息：响应
type EventVoteCreateResModel struct {
	Type		int8	`json:"type" form:"type"`
	EID			string	`json:"eid" form:"eid"`
}


// EventVoteUpdateReqModel 修改投票消息：请求
type EventVoteUpdateReqModel struct {
	EID				string		`json:"eid" form:"eid"`
	Title       	string   	`json:"title" form:"title"`
	Content       	string   	`json:"content" form:"content"`
	Tags          	[]string 	`json:"tags" form:"tags"`
	AllowedNumber 	int		   	`json:"allowed_number" form:"allowed_number"`
	Deadline      	string   	`json:"deadline" form:"deadline"`
	Options       	[]string 	`json:"options" form:"options"`
}

// EventVoteUpdateResModel 修改投票消息：响应
type EventVoteUpdateResModel struct {
	Frequency	int		`json:"frequency" form:"frequency"`
	EID			string	`json:"eid" form:"eid"`
}
