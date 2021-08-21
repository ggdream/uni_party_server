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
	UID				uint		`json:"uid"`
	EID				string		`json:"eid"`
	Title			string		`json:"title" `
	Type			int8		`json:"type"`
	CreateTime		string		`json:"create_time"`
	IsGet			bool		`json:"is_get"`
	IsCollect		bool		`json:"is_collect"`
}
