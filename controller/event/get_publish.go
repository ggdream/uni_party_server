package event

import (
	"gateway/cache"
	"gateway/middleware"
	"gateway/model/event"
	"gateway/mongo"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// PublicationGetController 获取用户发布的消息
func PublicationGetController(c *gin.Context) {
	var form event.GetPublicationReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 按offset和number去Redis中取出自己发布的消息
	ca := cache.Event{}
	eidList, err := ca.GetPublish(c.GetUint(middleware.KeyUID), form.Offset, form.Number)
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
