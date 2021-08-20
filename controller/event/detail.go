package event

import (
	"gateway/model/event"
	"gateway/mongo"
	"gateway/tools/errno"

	"github.com/gin-gonic/gin"
)

// DetailController 获取消息详情信息
func DetailController(c *gin.Context) {
	var form event.DetailReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeAccessTokenParsingErr)
		return
	}

	document := mongo.EventDocument{}
	data, err := document.FindOneDetail(form.EID)
	if err != nil {
		errno.Abort(c, errno.TypeMongoErr)
		return
	}

	errno.Perfect(c, data)
}
