package video

import (
	"gateway/cache"
	"gateway/middleware"
	"gateway/model/video"
	"gateway/sql"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
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

	// 分页查询用户收藏的视频vid
	ca := cache.Video{}
	vidList, err := ca.GetCollect(c.GetUint(middleware.KeyUID), form.Offset, form.Number)
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	// 去MySQL查询视频的详细数据
	db := sql.VideoInfoTable{}
	data, err := db.QueryIn(c.GetUint(middleware.KeyUID), vidList)
	if err != nil {
		errno.Abort(c, errno.TypeMySQLErr)
		return
	}
	// TODO: 按照Redis获取的列表顺序进行排序

	errno.Perfect(c, data)
}
