package middleware

import "github.com/gin-gonic/gin"


// TokenVerifyMW 从请求头获取token，并读取Redis进行验证
func TokenVerifyMW() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
