package user


// UserSettingsThemeGetReqModel 获取当前主题色：请求
type UserSettingsThemeGetReqModel struct {}

// UserSettingsThemeGetResModel 获取当前主题色：响应
type UserSettingsThemeGetResModel struct {
	Color		string	`json:"color" form:"color"`
}


// UserSettingsThemeSetReqModel 设置当前主题色：请求
type UserSettingsThemeSetReqModel struct {
	Color		string	`json:"color" form:"color"`
}

// UserSettingsThemeSetResModel 设置当前主题色：响应
type UserSettingsThemeSetResModel struct {}


// UserSettingsPushReqModel 获取推送选择项：请求
type UserSettingsPushReqModel struct {}


// UserSettingsPushResModel 获取推送选择项：响应
type UserSettingsPushResModel struct {
	Event	bool	`json:"event" form:"event"`
	Video	bool	`json:"video" form:"video"`
}


// UserSettingsPushEventReqModel 推送消息选择项：请求
type UserSettingsPushEventReqModel struct {
	Status	bool	`json:"status" form:"status"`
}

// UserSettingsPushEventResModel 推送消息选择项：响应
type UserSettingsPushEventResModel struct {}


// UserSettingsPushEventReqModel 推送视频选择项：请求
type UserSettingsPushVideoReqModel struct {
	Status	bool	`json:"status" form:"status"`
}

// UserSettingsPushEventResModel 推送视频选择项：响应
type UserSettingsPushVideoResModel struct {}
