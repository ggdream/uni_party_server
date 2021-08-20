package event

import (
	"gateway/cache"
	"gateway/middleware"
	"gateway/model/event"
	"gateway/tools/errno"

	"github.com/gin-gonic/gin"
)

// ArchiveAddAttendController 用户关注消息
func ArchiveAddAttendController(c *gin.Context) {
	var form event.ArchiveAddAttendReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeAccessTokenParsingErr)
		return
	}

	ca := cache.Event{}
	if err := ca.DoAttend(c.GetUint(middleware.KeyUID), form.Eid); err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	errno.Perfect(c, nil)
}

// ArchiveDelAttendController 用户取关消息
func ArchiveDelAttendController(c *gin.Context) {
	var form event.ArchiveDelAttendReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeAccessTokenParsingErr)
		return
	}

	ca := cache.Event{}
	if err := ca.DoUnAttend(c.GetUint(middleware.KeyUID), form.Eid); err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	errno.Perfect(c, nil)
}
