package event

import (
	"gateway/es"
	"gateway/model/event"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// SearchController 搜索用户时间线内的消息
func SearchController(c *gin.Context) {
	var form event.SearchReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// TODO: 获取学生所属uid

	search := es.EventIndex{}
	total, data, err := search.QueryByTitle(20, form.Query, form.Offset, form.Number)
	if err != nil {
		errno.Abort(c, errno.TypeESErr)
		return
	}

	var match []event.EventResultModel
	for _, v := range data {
		m := event.EventResultModel{
			UID:        v.UID,
			EID:        v.EID,
			Title:      v.Title,
			Type:       v.Type,
			CreateTime: v.Datetime,
		}
		match = append(match, m)
	}

	ret := &event.SearchResModel{
		Total: total,
		Match: nil,
	}
	errno.Perfect(c, ret)
}
