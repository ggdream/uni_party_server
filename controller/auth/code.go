package auth

import (
	"gateway/cache"
	"gateway/model/auth"
	"gateway/tools/errno"
	"gateway/tools/random"
	"github.com/gin-gonic/gin"
)

// CodePhoneController 发送手机验证码
func CodePhoneController(c *gin.Context) {
	var form auth.LoginSendPhoneCodeReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	deviceCode, smsCode := genCode()
	sms := cache.SMS{}
	if err := sms.Save(form.Phone, deviceCode, smsCode); err != nil {
		errno.Abort(c, errno.TypeSMSReqNotExpire)
		return
	}

	errno.New(c, errno.TypePerfect, deviceCode)
}

// CodeEMailController 发送短信验证码
func CodeEMailController(c *gin.Context) {
	var form auth.LoginSendEMailCodeReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	deviceCode, smsCode := genCode()
	sms := cache.SMS{}
	if err := sms.Save(form.EMail, deviceCode, smsCode); err != nil {
		// 缓存验证码失败
		errno.Abort(c, errno.TypeSMSReqNotExpire)
		return
	}
	errno.New(c, errno.TypePerfect, deviceCode)
}

// genCode 生成设备id和消息验证码
func genCode() (string, string) {
	return random.NewDeviceID(), random.NewSMSCode()
}
