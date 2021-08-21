package event

import (
	"gateway/cache"
	"gateway/middleware"
	"gateway/model/event"
	"gateway/mongo"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
	"time"
)

// ParticipationCreateController 发布报名类消息
func ParticipationCreateController(c *gin.Context) {
	var form event.ParticipationCreateReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 	消息发布至MongoDB，得到eid
	document := mongo.EventDocument{
		UID:      c.GetUint(middleware.KeyUID),
		Type:     TypeParticipation,
		Title:    form.Title,
		Content:  form.Content,
		Tags:     form.Tags,
		Datetime: time.Now(),
		Constraint: mongo.ParticipationField{
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

// ParticipationUpdateController 修改报名类消息
func ParticipationUpdateController(c *gin.Context) {
	var form event.ParticipationUpdateReqModel
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
		Constraint: mongo.ParticipationField{
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

// ParticipationJoinController 申请 || 取消 参加活动
func ParticipationJoinController(c *gin.Context) {
	var form event.ParticipationJoinReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// TODO: 查询该消息的具体元数据，是否允许撤销、截止日期，最多人数

	uid := c.GetUint(middleware.KeyUID)
	ca := cache.Event{}

	var err error
	if form.Type == 0 {
		// 参加
		err = ca.JoinParticipation(uid, form.EID, 50)
	} else if form.Type == 1 {
		// 取消
		err = ca.UnJoinParticipation(uid, form.EID)
	} else {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	errno.Perfect(c, nil)
}
