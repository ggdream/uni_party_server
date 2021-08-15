package video

import (
	"gateway/model/video"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// PushController 视频推送
func PushController(c *gin.Context) {
	var form video.PushReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// TODO: 获取20条视频数据
}
