package event


// userInfo 用户信息
type userInfo struct {
	UID				string		`json:"uid" form:"uid"`
	UName			string		`json:"uname" form:"uname"`
	AvatarURL		string		`json:"avatar_url" form:"avatar_url"`
	Motto			string		`json:"motto" form:"motto"`
	Sex				string		`json:"sex" form:"sex"`
	Type			int8		`json:"type" form:"type"`
	OrgName			string		`json:"org_name" form:"org_name"`
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
	UserInfo		userInfo	`json:"userinfo" form:"userinfo"`
	IsGet			bool		`json:"is_get" form:"is_get"`
	IsCollect		bool		`json:"is_collect" form:"is_collect"`
}
