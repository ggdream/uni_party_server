package event

// userInfo 用户信息
type userInfo struct {
	UID       string `json:"uid" form:"uid"`
	UName     string `json:"uname" form:"uname"`
	AvatarURL string `json:"avatar_url" form:"avatar_url"`
	Motto     string `json:"motto" form:"motto"`
	Sex       string `json:"sex" form:"sex"`
	Type      int8   `json:"type" form:"type"`
	OrgName   string `json:"org_name" form:"org_name"`
}

// EventResult 订阅消息返回的单条消息（模糊，不包括消息内容）
type EventResult struct {
	UID      uint   `json:"uid"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`

	EID        string `json:"eid"`
	Title      string `json:"title" `
	Type       int8   `json:"type"`
	CreateTime int64  `json:"create_time"`
	IsGet      bool   `json:"is_get"`
	IsAttend   bool   `json:"is_attend"`
}
