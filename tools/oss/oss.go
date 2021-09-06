package oss

import (
	"context"
	"errors"
	"io"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const bucket = "pan"

// Manager OSS管家
// 上传单文件，上传多文件，删除文件/文件夹，分享文件/文件夹，获取目录列表，获取文件数据
type Manager struct {
	Client  *minio.Client
	Context context.Context
}

func New() (*Manager, error) {
	options := minio.Options{
		Creds:  credentials.NewStaticV4("", "", ""),
		Secure: false,
	}
	client, err := minio.New("127.0.0.1:8080", &options)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
	defer cancel()

	found, err := client.BucketExists(timeoutCtx, bucket)
	if err != nil {
		return nil, err
	}
	if !found {
		bucketOption := minio.MakeBucketOptions{
			Region:        "cn-north-1",
			ObjectLocking: true,
		}
		err = client.MakeBucket(timeoutCtx, bucket, bucketOption)
		if err != nil {
			return nil, err
		}
	}

	return &Manager{
		Client:  client,
		Context: ctx,
	}, nil
}

// UploadSingleFile 上传单文件
func (m *Manager) UploadSingleFile(reader io.Reader, objectSize int64, objectName, uniqueId string) (minio.UploadInfo, error) {
	t := time.Duration(objectSize / 1024 / 1) // 1s上传1M
	ctx, cancel := context.WithTimeout(m.Context, t)
	defer cancel()

	option := minio.PutObjectOptions{
		UserTags: map[string]string{
			"name": objectName,
			"ext":  filepath.Ext(objectName),
		},
	}
	return m.Client.PutObject(ctx, bucket, uniqueId, reader, objectSize, option)
}

// GetObject 获取文件流
func (m *Manager) GetObject(objectName string, offset int64) (io.Reader, error) {
	ctx, cancel := context.WithTimeout(m.Context, 3*time.Second)
	defer cancel()

	object, err := m.Client.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	_, err = object.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	return object, nil
}

// ListObjects 获取路径下的文件
func (m *Manager) ListObjects(prefix string) ([]*Model, error) {
	ctx, cancel := context.WithTimeout(m.Context, 3*time.Second)
	defer cancel()

	option := minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: false,
	}
	objectsInfo := m.Client.ListObjects(ctx, bucket, option)

	data := make([]*Model, 0)
	for objectInfo := range objectsInfo {
		if objectInfo.Err != nil {
			return nil, objectInfo.Err
		}

		model := Model{
			Name: objectInfo.Key,
			Hash: objectInfo.ETag,
			Size: objectInfo.Size,
			Type: objectInfo.UserTags["type"],

			ContentType:  objectInfo.ContentType,
			LastModified: objectInfo.LastModified.Unix(),
		}
		data = append(data, &model)
	}

	return data, nil
}

// DelObjects 删除文件
func (m *Manager) DelObjects(modelList []DelModel) ([]string, error) {
	errObjects := make([]string, 0)

	for _, model := range modelList {
		ctx, cancel := context.WithTimeout(m.Context, 3*time.Second)
		defer cancel()

		if model.IsBlob {
			err := m.Client.RemoveObject(ctx, bucket, model.Prefix, minio.RemoveObjectOptions{})
			if err != nil {
				return nil, err
			}
			continue
		}

		option := minio.ListObjectsOptions{
			Prefix:    model.Prefix,
			Recursive: true,
		}
		objectsInfo := m.Client.ListObjects(ctx, bucket, option)
		errChan := m.Client.RemoveObjects(ctx, bucket, objectsInfo, minio.RemoveObjectsOptions{})

		for err := range errChan {
			if err.Err != nil {
				errObjects = append(errObjects, err.ObjectName)
			}
		}
	}

	if len(errObjects) != 0 {
		return errObjects, errors.New("del some objects failed")
	}
	return nil, nil
}
