package middleware

import (
	"gateway/tools/log"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	logger := log.New()
	logger.SetOutput(os.Stdout)

	return func(c *gin.Context) {
		visitTime := time.Now()
		visitUrl := c.Request.RequestURI
		userIp := c.ClientIP()

		c.Next()

		spentTime := time.Now().Sub(visitTime)
		fieldsMap := map[string]interface{}{
			"ip":  userIp,
			"sub": spentTime,
			"url": visitUrl,
		}
		logger.WithFields(fieldsMap).Info()
	}
}
