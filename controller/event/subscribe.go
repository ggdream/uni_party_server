package event

import (
	"gateway/model/event"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// SubscribeController 搜索用户时间线内的消息
func SubscribeController(c *gin.Context) {
	var form event.SubscribeReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// TODO: 查看用户所属的班级uid

	// TODO: 用班级uid查询MongoDB
}
