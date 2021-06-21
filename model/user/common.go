package user


// userInfo 用户基本信息
type userInfo struct {
	UID     string `json:"uid"`
	Uname   string `json:"uname"`
	Avatar  string `json:"avatar"`
	Sex     string `json:"sex"`
	Motto   string `json:"motto"`
	Type    int8   `json:"type"`
	OrgName string `json:"org_name"`
}
