package user

import (
	"gateway/es"
	"gateway/model/user"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

const pageSize = 20

// SearchController 搜索用户
func SearchController(c *gin.Context) {
	var form user.SearchReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}
	if form.Page < 1 {
		errno.Abort(c, errno.TypeParamsInvalidErr)
		return
	}

	search := es.UserIndex{}
	total, data, err := search.Query(form.Query, (form.Page-1)*pageSize, pageSize)
	if err != nil {
		errno.Abort(c, errno.TypeESErr)
		return
	}

	var match []user.SimpleUserInfoModel
	for _, v := range data {
		temp := user.SimpleUserInfoModel{
			UID:      v.UID,
			Username: v.Username,
			Avatar:   v.Avatar,
			College:  v.College,
		}
		match = append(match, temp)
	}

	ret := &user.SearchResModel{
		Total: total,
		Match: match,
	}
	errno.Perfect(c, ret)
}
