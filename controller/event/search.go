package event

import (
	"gateway/model/event"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// SearchController 搜索用户时间线内的消息
func SearchController(c *gin.Context) {
	var form event.SearchReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// TODO: 去ElasticSearch里查询
}
