package helper

import (
	"gateway/cache"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// ContactController 联系我们的方式
func ContactController(c *gin.Context) {
	ca := cache.Contact{}
	data, err := ca.GetInfo()
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	errno.Perfect(c, data)
}
