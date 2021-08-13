package helper

import (
	"gateway/model/helper"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// FeedbackController 接收用户反馈
func FeedbackController(c *gin.Context) {
	var form helper.FeedbackReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// TODO: 记录到MongoDB
	errno.Perfect(c, nil)
}
