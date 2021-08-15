package file

import (
	"context"
	"mime/multipart"
	"time"
)

// Upload 上传文件*multipart.FileHeader 到七牛云OSS指定路径
func Upload(file *multipart.FileHeader, dest string) error {
	fileReader, err := file.Open()
	if err != nil {
		return err
	}

	fileHandler := New("atLchdSy60cV5zsWf5Mha3FqSxyP1ui40iWQ3VFc", "4C0EjtYwmzO07SPaWRiolYV8519vwY1UCYEGfix4")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return fileHandler.UploadBySteam(ctx, fileReader, file.Size, dest)
}
