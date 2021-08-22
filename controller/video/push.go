package video

import (
	"gateway/model/video"
	"gateway/sql"
	"gateway/tools/crypto"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
	"strconv"
	"unsafe"
)

// PushController 视频推送
func PushController(c *gin.Context) {
	var form video.PushReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	aes, err := crypto.AES([]byte("0102030405060708"), []byte("0102030405060708"))
	if err != nil {
		errno.Abort(c, errno.TypeUnknownMistake)
		return
	}
	plaint, err := aes.Decrypt([]byte(form.Pointer), crypto.WrapperBase64)
	if err != nil {
		errno.Abort(c, errno.TypeDecryptFailed)
		return
	}
	id, err := strconv.Atoi(*(*string)(unsafe.Pointer(&plaint)))
	if err != nil {
		errno.Abort(c, errno.TypeParamsInvalidErr)
		return
	}

	db := sql.VideoInfoTable{}
	data, err := db.QueryPush(id, pageSize)
	if err != nil {
		errno.Abort(c, errno.TypeMySQLErr)
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
			CreateTime: v.CreatedAt.Unix(),
			Status:     v.Status,
		}
		result = append(result, res)
	}

	hasNext := false
	pointer := ""
	last := data[len(data)-1]
	if last.ID != 0 {
		hasNext = true
		byteId := []byte(strconv.Itoa(int(last.ID)))
		pointer = aes.Encrypt(byteId).ToBase64()
	}

	final := video.PushResModel{
		Pointer: pointer,
		HasNext: hasNext,
		Result:  result,
	}
	errno.Perfect(c, final)
}
