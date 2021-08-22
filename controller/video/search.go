package video

import (
	"gateway/es"
	"gateway/model/video"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// SearchController 搜索视频
func SearchController(c *gin.Context) {
	var form video.SearchReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}
	if form.Page < 1 {
		errno.Abort(c, errno.TypeParamsInvalidErr)
		return
	}

	search := es.VideoIndex{}
	total, data, err := search.QueryByTitle(form.Keyword, (form.Page-1)*pageSize, pageSize)
	if err != nil {
		errno.Abort(c, errno.TypeESErr)
		return
	}

	var result []video.VideoResultModel
	for _, v := range data {
		res := video.VideoResultModel{
			VID:        v.VID,
			UID:        v.UID,
			Title:      v.Title,
			Cover:      v.Cover,
			Tags:       v.Video,
			CreateTime: v.Datetime,
			Status:     v.Status,
		}
		result = append(result, res)
	}

	ret := &video.SearchResModel{
		Result: result,
		Total:  total,
	}
	errno.Perfect(c, ret)
}
