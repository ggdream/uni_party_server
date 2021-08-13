package event

import (
	"gateway/model/event"
	"gateway/mongo"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// DeleteController 删除用户发布的某条消息
func DeleteController(c *gin.Context) {
	var form event.DeleteReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 删除在MongoDB中的记录
	document := mongo.EventDocument{}
	err := document.Delete(form.EID)
	if err != nil {
		errno.Abort(c, errno.TypeUnknownMistake)
		return
	}

	errno.Perfect(c, nil)
}
