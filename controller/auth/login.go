package auth

import (
	"errors"
	"gateway/cache"
	"gateway/model/auth"
	"gateway/sql"
	"gateway/tools/errno"
	"gateway/tools/safety"
	"gateway/tools/ytoken"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var handler = ytoken.NewHandler([]byte("0123456789"))

func genSign(uid uint) (string, error) {
	constraint := ytoken.Constraint{
		Signer:      "uparty",
		Expiry:      time.Now().Add(30 * time.Minute).Unix(),
		Serial:      "",
		Beneficiary: strconv.Itoa(int(uid)),
	}
	yToken := ytoken.New(&constraint)
	return handler.Sign(yToken)
}

// LoginByCipherController 密码登录
func LoginByCipherController(c *gin.Context) {
	var form auth.LoginCipherReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 查询记录是否存在。不存在则返回未注册
	user := sql.UserInfoTable{}
	err := user.QueryUserByPhone(form.Account)
	if err != nil {
		if errors.Is(err, sql.ErrRecordNotFound) {
			// 没有这条记录
			errno.Abort(c, errno.TypeNotRegister)
		} else {
			errno.Abort(c, errno.TypeUnknownMistake)
		}
		return
	}

	// 若获取后匹配不成功，则返回密码有误
	sAuth := safety.ScryptAuth{
		Password: form.Password,
		Salt:     user.Salt,
	}
	isMatch, err := sAuth.Compare(user.Password)
	if err != nil {
		errno.Abort(c, errno.TypeCipherNotAllowed)
		return
	}
	if isMatch == false {
		errno.Abort(c, errno.TypeCipherNotMatch)
		return
	}

	// 若获取后匹配成功，则生成JWT返回客户端
	sign, err := genSign(user.ID)
	if err != nil {
		errno.Abort(c, errno.TypeUnknownMistake)
		return
	}

	errno.New(c, errno.TypePerfect, sign)
}


// LoginByCodeController 验证码登录
// 不区分登录与注册
func LoginByCodeController(c *gin.Context) {
	var form auth.LoginCodeReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 首先去Redis查找，key为从form中获取到的设备码，获取value与form的sms对比
	sms := cache.SMS{}
	ok, err := sms.Verify(form.Account, form.DeviceCode, form.SmsCode)
	if err != nil {
		// 	如果Redis没有相对应的设备码，则返回设备码错误
		errno.Abort(c, errno.TypeUnknownMistake)
		return
	}
	if !ok {
		// 	如果获取到的value与sms不匹配，则用户输入了错误的短信验证码
		errno.Abort(c, errno.TypeCodeVerifyFailed)
		return
	}

	// 	如果正确，则查看是否有该用户
	// 查询记录是否存在。不存在则返回未注册
	user := sql.UserInfoTable{}
	err = user.QueryUserByPhone(form.Account)
	if err != nil {
		if errors.Is(err, sql.ErrRecordNotFound) {
			// 创建新用户失败
			if err := user.Create(); err != nil {
				errno.Abort(c, errno.TypeUserCreateFailed)
				return
			}
		} else {
			// 查询错误
			errno.Abort(c, errno.TypeUnknownMistake)
			return
		}
	}

	// 	生成JWT返回客户端
	sign, err := genSign(user.ID)
	if err != nil {
		errno.Abort(c, errno.TypeUnknownMistake)
		return
	}

	errno.New(c, errno.TypePerfect, sign)
}
