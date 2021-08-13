package helper

import (
	"gateway/cache"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// UpdateController 获取应用版本、升级信息
func UpdateController(c *gin.Context)  {
	ca := cache.Application{}
	data, err := ca.GetInfo()
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	errno.Perfect(c, data)
}
