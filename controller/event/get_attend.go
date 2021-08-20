package event

import (
	"gateway/cache"
	"gateway/middleware"
	"gateway/model/event"
	"gateway/mongo"
	"gateway/tools/errno"

	"github.com/gin-gonic/gin"
)

// AttendGetController 获取用户关注的消息
func AttendGetController(c *gin.Context) {
	var form event.GetCollectionReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 按offset和number去Redis中取出自己关注的消息
	ca := cache.Event{}
	eidList, err := ca.GetAttend(c.GetUint(middleware.KeyUID), form.Offset, form.Number)
	if err != nil {
		if err == cache.MustGEZeroErr {
			errno.Abort(c, errno.TypeParamsParsingErr)
		} else {
			errno.Abort(c, errno.TypeCacheErr)
		}
		return
	}

	document := mongo.EventDocument{}
	result, err := document.FindIn(eidList)
	if err != nil {
		errno.Abort(c, errno.TypeUnknownMistake)
		return
	}

	errno.Perfect(c, result)
}
