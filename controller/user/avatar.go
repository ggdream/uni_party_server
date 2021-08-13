package user

import (
	"context"
	"gateway/middleware"
	"gateway/model/user"
	"gateway/tools/errno"
	"gateway/tools/file"
	"github.com/gin-gonic/gin"
	"time"
)

// UploadAvatarController 更新用户头像
func UploadAvatarController(c *gin.Context) {
	var form user.AvatarUploadReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	fileReader, err := form.File.Open()
	if err != nil {
		errno.Abort(c, errno.TypeFileOpenFailed)
		return
	}

	fileUploader := file.New("atLchdSy60cV5zsWf5Mha3FqSxyP1ui40iWQ3VFc", "4C0EjtYwmzO07SPaWRiolYV8519vwY1UCYEGfix4")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	err = fileUploader.UploadBySteam(ctx, fileReader, form.File.Size, "avatars/u/"+c.GetString(middleware.KeyUID))
	if err != nil {
		errno.Abort(c, errno.TypeFileUploadFailed)
		return
	}

	errno.Perfect(c, "avatars/u/"+c.GetString(middleware.KeyUID))
}
