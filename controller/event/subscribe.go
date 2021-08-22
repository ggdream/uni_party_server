package event

import (
	"gateway/cache"
	"gateway/middleware"
	"gateway/model/event"
	"gateway/mongo"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

const (
	pageSize = 20
)

// SubscribeController 搜索用户时间线内的消息
func SubscribeController(c *gin.Context) {
	var form event.SubscribeReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}
	if form.Page < 1 {
		errno.Abort(c, errno.TypeParamsInvalidErr)
		return
	}

	// TODO: 查看用户所属的班级uid
	classUid := uint(555)
	uid := c.GetUint(middleware.KeyUID)

	var size int64
	var counter int64
	ca := cache.Event{}
	if form.Page == 1 {
		var err error
		counter, err = ca.GetUnreadAndReset(classUid, uid)
		if err != nil {
			errno.Abort(c, errno.TypeCacheErr)
			return
		}

		if counter > pageSize {
			size = counter
		} else {
			size = pageSize
		}
	} else {
		size = pageSize
	}

	document := mongo.EventDocument{}
	data, err := document.Find(classUid, (form.Page-1)*size, size)
	if err != nil {
		errno.Abort(c, errno.TypeMongoErr)
		return
	}

	eidList := make([]interface{}, len(data))
	for i, v := range data {
		eidList[i] = v.EID
	}
	getList, err := ca.IsGet(uid, eidList)
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}
	attendList, err := ca.IsAttend(uid, eidList)
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	result := make([]event.EventResult, len(data))
	for i, v := range data {
		result[i] = event.EventResult{
			UID:        v.UID,
			Avatar:     v.Avatar,
			Username:   v.Username,
			EID:        v.EID,
			Title:      v.Title,
			Type:       v.Type,
			CreateTime: v.Datetime.Unix(),
			IsGet:      getList[i],
			IsAttend:   attendList[i],
		}
	}

	ret := &event.SubscribeResModel{
		Unread: counter,
		Result: result,
	}
	errno.Perfect(c, ret)
}
