package user


// UserEventPublicationsReqModel 获取用户发布的消息：请求
type UserEventPublicationsReqModel struct {
	UID    int `json:"uid" form:"uid"`
	Offset int `json:"offset" form:"offset"`
	Number int `json:"number" form:"number"`
}

// UserEventPublicationsResModel 获取用户发布的消息：响应
type UserEventPublicationsResModel struct {
	Total  int             `json:"total" form:"total"`
	Result []publishResult `json:"result" form:"result"`
}


// publishResult 用户发布的消息元数据（相比eventResult{}，没有UserInfo字段）
type publishResult struct {
	EID            string   `json:"eid" form:"eid"`
	Title          string   `json:"title" form:"title"`
	Type           int8     `json:"type" form:"type"`
	Tags           []string `json:"tags" form:"tags"`
	WatchCounter   int      `json:"watch_counter" form:"watch_counter"`
	StarCounter    int      `json:"star_counter" form:"star_counter"`
	CommentCounter int      `json:"comment_counter" form:"comment_counter"`
	CreateTime     string   `json:"create_time" form:"create_time"`
	UpdateTime     string   `json:"update_time" form:"update_time"`
	IsGet          bool     `json:"is_get" form:"is_get"`
	IsCollect      bool     `json:"is_collect" form:"is_collect"`
}



// UserEventCollectionsReqModel 获取用户收藏的消息：请求
type UserEventCollectionsReqModel struct {
	UID    int `json:"uid" form:"uid"`
	Offset int `json:"offset" form:"offset"`
	Number int `json:"number" form:"number"`
}

// UserEventCollectionsResModel 获取用户收藏的消息：响应
type UserEventCollectionsResModel struct {
	Total  int           `json:"total" form:"total"`
	Result []eventResult `json:"result" form:"result"`
}


// eventResult 订阅消息返回的单条消息（模糊，不包括消息内容）
type eventResult struct {
	EID				string		`json:"eid" form:"eid"`
	Title			string		`json:"title" form:"title"`
	Type			int8		`json:"type" form:"type"`
	Tags			[]string	`json:"tags" form:"tags"`
	CreateTime		string		`json:"create_time" form:"create_time"`
	UpdateTime		string		`json:"update_time" form:"update_time"`
	GetCounter		int			`json:"get_counter" form:"get_counter"`
	WatchCounter	int			`json:"watch_counter" form:"watch_counter"`
	CommentCounter	int			`json:"comment_counter" form:"comment_counter"`
	UserInfo		UserInfo	`json:"userinfo" form:"userinfo"`
	IsGet			bool		`json:"is_get" form:"is_get"`
	IsCollect		bool		`json:"is_collect" form:"is_collect"`
}
