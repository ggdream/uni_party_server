package user

// UserInfo 用户基本信息
type UserInfo struct {
	ID       string `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Sex      string `json:"sex"`
	Motto    string `json:"motto"`
	Type     int8   `json:"type"`
	OrgName  string `json:"org_name"`
}
