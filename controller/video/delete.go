package video

import (
	"gateway/model/video"
	"gateway/sql"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// DeleteController 删除视频
func DeleteController(c *gin.Context) {
	var form video.DeleteReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	db := sql.VideoInfoTable{}
	if err := db.Delete(form.Vid); err != nil {
		errno.Abort(c, errno.TypeVideoDeleteFailed)
		return
	}

	errno.Perfect(c, nil)
}
