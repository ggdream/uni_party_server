package video

import (
	"gateway/cache"
	"gateway/middleware"
	"gateway/model/video"
	"gateway/sql"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	pageSize = 20
)

// CollectController 用户收藏视频
func CollectController(c *gin.Context) {
	var form video.CollectReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	ca := cache.Video{}
	err := ca.AddCollect(c.GetUint(middleware.KeyUID), form.Vid)
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	errno.Perfect(c, nil)
}

// CollectGetController 获取用户收藏的视频
func CollectGetController(c *gin.Context) {
	var form video.CollectGetReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}
	if form.Page < 1 {
		errno.Abort(c, errno.TypeParamsInvalidErr)
		return
	}

	// 分页查询用户收藏的视频vid
	ca := cache.Video{}
	total, err := ca.CountCollect(c.GetUint(middleware.KeyUID))
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}
	vidAndScoreList, err := ca.GetCollect(c.GetUint(middleware.KeyUID), (form.Page-1)*pageSize, pageSize)
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	vidList := make([]string, len(vidAndScoreList)/2)
	scoreList := make([]string, len(vidAndScoreList)/2)

	// 去MySQL查询视频的详细数据
	db := sql.VideoInfoTable{}
	data, err := db.QueryIn(vidList)
	if err != nil {
		errno.Abort(c, errno.TypeMySQLErr)
		return
	}

	var result []video.VideoResultModel
	for i, v := range data {
		no, _ := strconv.ParseInt(scoreList[i], 10, 64)
		res := video.VideoResultModel{
			VID:         v.VID,
			UID:         v.UID,
			Title:       v.Title,
			Cover:       v.Cover,
			Tags:        v.Video,
			CollectTime: no,
			CreateTime:  v.CreatedAt.Unix(),
			Status:      v.Status,
		}
		result = append(result, res)
	}

	final := video.CollectGetResModel{
		Total:  total,
		Result: result,
	}
	errno.Perfect(c, final)
}
