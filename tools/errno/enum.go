package errno

const (
	// 正常
	TypePerfect = Type(iota)

	// 万能错误
	TypeUnknownMistake

	// 权限类
	TypeNotLoggedIn           // 未登录
	TypeNotRegister           // 未注册
	TypeLoggedInFailed        // 登录失败
	TypeSMSReqNotExpire       // 验证码请求未过期限
	TypeUnauthorizedAccess    // 非法访问
	TypeCodeVerifyFailed      // 短信验证失败
	TypeCipherNotAllowed      // 不允许的密码
	TypeCipherNotMatch        // 密码错误
	TypeAccessTokenParsingErr // 登陆凭据解析错误
	TypeAccessTokenExpired    // 登陆凭据过期
	TypeRefreshTokenExpired   // 刷新凭据过期

	TypeUserCreateFailed // 用户创建失败

	TypeFileOpenFailed   // 文件打开失败
	TypeFileUploadFailed // 文件上传失败

	TypeEventPublishFailed // 消息发布失败
	TypeEventUpdateFailed  // 消息更新失败

	TypeVideoDeleteFailed // 视频删除失败

	// 参数类
	TypeParamsMissingErr // 参数缺失错误
	TypeParamsParsingErr // 参数解析错误
	TypeParamsInvalidErr // 参数无效错误

	// 安全类
	TypeServerIsBusy         // 服务器忙
	TypeVisitTooFast         // 访问过快
	TypeEncKeyGetFailed      // 获取加密秘钥失败
	TypeCryptoInstanceFailed // 实例化加解密器失败
	TypeDecryptFailed        // 解密失败

	// 后台类
	TypeCacheErr // Redis缓存出错
	TypeMongoErr // MongoDB出错
	TypeMySQLErr // MySQL出错

	TypeEventErr              // 消息类错误
	TypeEventTypeErr          // 消息类型错误
	TypeEventDeadlineErr      // 已到期限错误
	TypeEventOptionsNumberErr // 用户投票选项数量超过限值
	TypeEventVoteAOErr        // 答案个数与选项个数不符
	TypeEventCannotCancelErr  // 不能取消
	TypeEventQuotaFullErr     // 名额已满

	TypeESErr // ElasticSearch出错
)

type Type int16

func (t Type) Index() int16 {
	return int16(t)
}

func (t Type) String() string {
	return getErrFromMap(t)
}
