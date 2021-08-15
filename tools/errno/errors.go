package errno

// getErrFromMap 从定义的errorsMap中获取文字描述信息
func getErrFromMap(t Type) string {
	return errorsMap[t]
}

var errorsMap = map[Type]string{
	TypePerfect: "",

	TypeUnknownMistake: "发生错误了呦~",

	TypeNotLoggedIn:           "用户未登录",
	TypeNotRegister:           "用户未注册",
	TypeLoggedInFailed:        "登录失败",
	TypeSMSReqNotExpire:       "验证码请求未过期限",
	TypeUnauthorizedAccess:    "非法访问",
	TypeCodeVerifyFailed:      "短信验证失败",
	TypeCipherNotMatch:        "账号或密码有误",
	TypeAccessTokenParsingErr: "登陆凭据解析失败",
	TypeAccessTokenExpired:    "登陆凭据过期",
	TypeRefreshTokenExpired:   "刷新凭据过期",

	TypeUserCreateFailed: "用户创建失败",

	TypeFileOpenFailed:   "文件打开失败",
	TypeFileUploadFailed: "文件上传失败",

	TypeEventPublishFailed: "消息发布失败",
	TypeEventUpdateFailed:  "消息更新失败",

	TypeVideoDeleteFailed: "视频删除失败",

	TypeParamsMissingErr: "参数缺失错误",
	TypeParamsParsingErr: "参数解析错误",

	TypeServerIsBusy: "服务器忙",
	TypeVisitTooFast: "访问过快",

	TypeCacheErr: "服务器有点忙",
	TypeMongoErr: "服务器有点忙",
	TypeMySQLErr: "服务器有点忙",
}
