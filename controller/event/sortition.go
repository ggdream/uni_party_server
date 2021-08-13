package event

import (
	"gateway/model/event"
	"gateway/mongo"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
	"time"
)

// SortitionCreateController 发布随机类消息
func SortitionCreateController(c *gin.Context) {
	var form event.SortitionCreateReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 	消息发布至MongoDB，得到eid
	document := mongo.EventDocument{
		UID:      c.GetUint("uid"),
		Type:     TypeSortition,
		Title:    form.Title,
		Content:  form.Content,
		Tags:     form.Tags,
		Datetime: time.Now(),
		Constraint: mongo.SortitionField{
			RequiredNumber: form.RequiredNumber,
			AllowedCancel:  form.AllowedCancel,
			Deadline:       form.Deadline,
		},
	}
	if err := document.Insert(); err != nil {
		errno.Abort(c, errno.TypeEventPublishFailed)
		return
	}

	// 返回结果
	errno.Perfect(c, document)
}

// SortitionUpdateController 修改随机类消息
func SortitionUpdateController(c *gin.Context) {
	var form event.SortitionUpdateReqModel
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
		Constraint: mongo.SortitionField{
			RequiredNumber: form.RequiredNumber,
			AllowedCancel:  form.AllowedCancel,
			Deadline:       form.Deadline,
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
