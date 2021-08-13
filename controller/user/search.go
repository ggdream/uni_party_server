package user

import (
	"gateway/model/user"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// SearchController 搜索用户
func SearchController(c *gin.Context) {
	var form user.SearchReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// TODO: 从ElasticSearch里查询
}
