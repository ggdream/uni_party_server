package video

import (
	"fmt"
	"gateway/cache"
	"gateway/middleware"
	"gateway/model/video"
	"gateway/sql"
	"gateway/tools/errno"
	"gateway/tools/file"
	"github.com/gin-gonic/gin"
)

const (
	// StatusUnreviewed 未审核
	StatusUnreviewed = iota
	// StatusDidNotPass 未通过
	StatusDidNotPass
	// StatusNormal 正常
	StatusNormal
	// StatusLocking 锁定
	StatusLocking
	// StatusDeletedByPlatform 被平台删除
	StatusDeletedByPlatform
	// StatusDeletedByUser 被用户删除
	StatusDeletedByUser
)

// UploadController 上传视频
func UploadController(c *gin.Context) {
	var form video.UploadReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 上传封面和视频到七牛云
	coverPath := fmt.Sprintf("archive/covers/%s.jpg", form.Cover.Filename)
	err := file.Upload(form.Cover, coverPath)
	if err != nil {
		errno.Abort(c, errno.TypeFileUploadFailed)
		return
	}
	videoPath := fmt.Sprintf("archive/videos/%s.mp4", form.Video.Filename)
	err = file.Upload(form.Video, videoPath)
	if err != nil {
		errno.Abort(c, errno.TypeFileUploadFailed)
		return
	}

	// 记录到MySQL里
	db := sql.VideoInfoTable{
		UID:      c.GetUint(middleware.KeyUID),
		VID:      "",
		Title:    form.Title,
		Tags:     form.Tags,
		Cover:    coverPath,
		Video:    coverPath,
		Location: form.Location,
		Status:   StatusUnreviewed,
	}
	if err := db.Create(); err != nil {
		errno.Abort(c, errno.TypeMySQLErr)
		return
	}

	// TODO: 写入Kafka消息队列，等待审核
}

// PublishGetController 获取用户发布的视频
func PublishGetController(c *gin.Context) {
	var form video.UploadGetReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 去MySQL分页查询用户发布的视频
	db := sql.VideoInfoTable{}
	data, err := db.QueryPage(form.UID, form.Offset, form.Number)
	if err != nil {
		errno.Abort(c, errno.TypeMySQLErr)
		return
	}

	// 去Redis查询用户一共发布的视频数量
	ca := cache.Video{}
	uploadCounter, err := ca.CountUpload(form.UID)
	if err != nil {
		errno.Abort(c, errno.TypeCacheErr)
		return
	}

	result := video.UploadGetResModel{
		Total:  uploadCounter,
		Result: make([]video.VideoResultLess, 0, len(data)),
	}
	for _, videoInfoTable := range data {
		temp := video.VideoResultLess{
			Vid:            videoInfoTable.VID,
			Title:          videoInfoTable.Title,
			Cover:          videoInfoTable.Cover,
			Tags:           videoInfoTable.Tags,
			CreateTime:     videoInfoTable.CreatedAt.Unix(),
			WatchCounter:   "",
			StarCounter:    "",
			CommentCounter: "",
			IsGet:          "",
			IsCollect:      "",
		}
		result.Result = append(result.Result, temp)
	}

	errno.Perfect(c, result)
}
