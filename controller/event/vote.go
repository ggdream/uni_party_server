package event

import (
	"gateway/model/event"
	"gateway/mongo"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
	"time"
)

// VoteCreateController 发布投票类消息
func VoteCreateController(c *gin.Context) {
	var form event.VoteCreateReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 	消息发布至MongoDB，得到eid
	document := mongo.EventDocument{
		UID:      c.GetUint("uid"),
		Type:     TypeVote,
		Title:    form.Title,
		Content:  form.Content,
		Tags:     form.Tags,
		Datetime: time.Now(),
		Constraint: mongo.VoteField{
			AllowedNumber: form.AllowedNumber,
			Deadline:      form.Deadline,
			Options:       form.Options,
		},
	}
	if err := document.Insert(); err != nil {
		errno.Abort(c, errno.TypeEventPublishFailed)
		return
	}

	// 返回结果
	errno.Perfect(c, document)
}

// VoteUpdateController 修改投票类消息
func VoteUpdateController(c *gin.Context) {
	var form event.VoteUpdateReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 对MongoDB进行update操作
	document := mongo.EventDocument{}
	updateDocument := mongo.UpdateEventDocument{
		Title:   form.Title,
		Content: form.Content,
		Tags:    form.Tags,
		Constraint: mongo.VoteField{
			AllowedNumber: form.AllowedNumber,
			Deadline:      form.Deadline,
			Options:       form.Options,
		},
	}
	_, err := document.Update(form.EID, &updateDocument)
	if err != nil {
		errno.Abort(c, errno.TypeEventUpdateFailed)
		return
	}

	// 返回结果
	errno.Perfect(c, updateDocument)
}
