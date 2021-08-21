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

// SortitionJoinController 申请 || 取消 参加随机
func SortitionJoinController(c *gin.Context) {
	var form event.SortitionJoinReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// TODO: 查询该消息的具体元数据，是否允许撤销、截止日期，最多人数

	uid := c.GetUint(middleware.KeyUID)
	ca := cache.Event{}

	var err error
	var res int64
	if form.Type == 0 {
		// 参加
		res, err = ca.JoinSortition(uid, form.EID)
	} else if form.Type == 1 {
		// 取消
		res, err = ca.UnJoinSortition(uid, form.EID)
	} else {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}
	if res == 0 {
		errno.New(c, errno.TypeEventErr, nil, "你已参加或取消该活动")
		return
	}

	errno.Perfect(c, nil)
}
