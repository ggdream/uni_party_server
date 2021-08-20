package event

import (
	"gateway/model/event"
	"gateway/sql"
	"gateway/tools/errno"

	"github.com/gin-gonic/gin"
)

// ArchiveGetController 用户进行get操作（自动、手动）
func ArchiveGetController(c *gin.Context) {
	var form event.ArchiveGetReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeAccessTokenParsingErr)
		return
	}

	// 插入记录
	db := sql.EventGetTable{}
	if err := db.Create(); err != nil {
		errno.Abort(c, errno.TypeMySQLErr)
		return
	}

	errno.Perfect(c, nil)
}
