package video

import (
	"gateway/model/video"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// SearchController 搜索视频
func SearchController(c *gin.Context)  {
	var form video.SearchReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// TODO: 去ElasticSearch里查询
}
