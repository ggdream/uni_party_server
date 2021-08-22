package helper

import (
	"gateway/middleware"
	"gateway/model/helper"
	"gateway/mongo"
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

	document := mongo.FeedbackDocument{
		UID:     c.GetUint(middleware.KeyUID),
		Email:   form.Email,
		Type:    form.Type,
		Title:   form.Title,
		Content: form.Content,
	}
	err := document.Insert()
	if err != nil {
		errno.Abort(c, errno.TypeMongoErr)
		return
	}

	errno.Perfect(c, nil)
}
