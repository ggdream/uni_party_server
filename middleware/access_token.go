package middleware

import (
	"gateway/model/auth"
	"gateway/tools/errno"
	"gateway/tools/ytoken"
	"github.com/gin-gonic/gin"
)

// AccessToken 从请求头获取token，进行验证
func AccessToken() gin.HandlerFunc {
	key := []byte("0123456789")
	tokenHandler := ytoken.NewHandler(key)

	return func(c *gin.Context) {
		var header auth.AccessTokenHeaderModel
		// 没有携带token
		if err := c.ShouldBindHeader(&header); err != nil {
			errno.Abort(c, errno.TypeNotLoggedIn)
			return
		}

		// token解析失败或者token内容不匹配
		yToken, isEqual, err := tokenHandler.Verify(header.Token)
		if err != nil || !isEqual {
			errno.Abort(c, errno.TypeAccessTokenExpired)
			return
		}

		// token已超时失效
		if yToken.Timeout() {
			errno.Abort(c, errno.TypeAccessTokenExpired)
			return
		}

		c.Set(KeyUID, yToken.Constraint.Beneficiary)
		c.Next()
	}
}
