package file

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

const (
	bucket = "uni-static"
)

type QiNiu struct {
	accessKey string
	secretKey string
	bucket    string
	formUploader *storage.FormUploader
}

func New(accessKey, secretKey string) *QiNiu {
	qiniu := &QiNiu{
		accessKey: accessKey,
		secretKey: secretKey,
		bucket:    bucket,
	}
	qiniu.genFormUploader()
	return qiniu
}

func (q *QiNiu) genToken() string {
	putPolicy := storage.PutPolicy{
		Scope: q.bucket,
	}
	mac := qbox.NewMac(q.accessKey, q.secretKey)
	return putPolicy.UploadToken(mac)
}

func (q *QiNiu) genFormUploader() {
	cfg := storage.Config{
		Zone:     &storage.ZoneHuanan,
		UseHTTPS: true,
	}
	q.formUploader = storage.NewFormUploader(&cfg)
}

// UploadBySteam 以流方式上传文件
func (q *QiNiu) UploadBySteam(context context.Context, file io.Reader, size int64, filepath string) error {
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	return q.formUploader.Put(context, &ret, q.genToken(), filepath, file, size, &putExtra)
}
