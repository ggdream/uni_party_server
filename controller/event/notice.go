package event

import (
	"context"
	"gateway/model/event"
	"gateway/mongo"
	"gateway/tools/errno"
	"gateway/tools/file"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"sync"
	"time"
)

// NoticeCreateController 发布通知类消息
func NoticeCreateController(c *gin.Context) {
	var form event.NoticeCreateReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 将文件保存至七牛云
	defer func() {
		if err := recover(); err != nil {
			errno.Abort(c, errno.TypeUnknownMistake)
		}
	}()

	var wg sync.WaitGroup
	var lock sync.Mutex
	ctx := context.Background()
	fileStore := make([]*mongo.FileField, 0)
	fileHandler := file.New("atLchdSy60cV5zsWf5Mha3FqSxyP1ui40iWQ3VFc", "4C0EjtYwmzO07SPaWRiolYV8519vwY1UCYEGfix4")

	for _, header := range form.Files {
		wg.Add(1)
		go func(file multipart.FileHeader) {
			defer wg.Done()
			fileReader, err := file.Open()
			if err != nil {
				panic(errno.TypeFileOpenFailed)
			}
			defer fileReader.Close()

			timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			filename := file.Filename
			fileRecord := mongo.FileField{
				Name: file.Filename,
				Url:  "",
			}
			lock.Lock()
			fileStore = append(fileStore, &fileRecord)
			lock.Unlock()

			err = fileHandler.UploadBySteam(timeoutCtx, fileReader, file.Size, filename)
			if err != nil {
				panic(errno.TypeFileUploadFailed)
			}
		}(*header)
	}
	wg.Wait()

	// TODO: 计算出所选择的class uid，记录到redis里
	//ca := cache.Event{}
	//if err := ca.AddUnread(c.GetUint(middleware.KeyUID)); err != nil {
	//
	//}

	// 	消息发布至MongoDB，得到eid
	document := mongo.EventDocument{
		UID:        c.GetUint("uid"),
		Type:       TypeNotice,
		Title:      form.Title,
		Content:    form.Content,
		Datetime:   time.Now(),
		Constraint: mongo.NoticeField{Files: fileStore},
	}
	if err := document.Insert(); err != nil {
		errno.Abort(c, errno.TypeEventPublishFailed)
		return
	}

	// 返回结果
	errno.Perfect(c, document)
}

// NoticeUpdateController 修改通知类消息
func NoticeUpdateController(c *gin.Context) {
	var form event.NoticeUpdateReqModel
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
	}
	_, err := document.Update(form.EID, &updateDocument)
	if err != nil {
		errno.Abort(c, errno.TypeEventUpdateFailed)
		return
	}

	// 返回结果
	errno.Perfect(c, updateDocument)
}
