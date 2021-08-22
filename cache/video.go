package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// Video 视频
// 用户上传使用：List；用户收藏使用：List
type Video struct{}

// joinUserUploadKey 获取用户上传视频列表的string-key
func (v Video) joinUserUploadKey(uid uint) string {
	return fmt.Sprintf("video:user:upload:%d", uid)
}

// joinUserCollectKey 获取用户收藏视频列表的string-key
func (v Video) joinUserCollectKey(uid uint) string {
	return fmt.Sprintf("video:user:collect:%d", uid)
}

// AddUpload 用户上传视频
func (v Video) AddUpload(uid uint, vid string) error {
	return client.LPush(v.joinUserUploadKey(uid), vid)
}

// AddCollect 用户收藏视频
func (v Video) AddCollect(uid uint, vid string) error {
	value := &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: vid,
	}
	return client.ZAdd(v.joinUserCollectKey(uid), value)
}

// DelUpload 用户删除上传视频
func (v Video) DelUpload(uid uint, vid string) error {
	return client.LRem(v.joinUserUploadKey(uid), 0, vid)
}

// DelCollect 用户删除收藏视频
func (v Video) DelCollect(uid uint, vid string) error {
	_, err := client.ZRem(v.joinUserCollectKey(uid), vid)
	return err
}

// CountUpload 获取用户上传视频的数量
func (v Video) CountUpload(uid uint) (int64, error) {
	return client.LLen(v.joinUserUploadKey(uid))
}

// CountCollect 获取用户收藏视频的数量
func (v Video) CountCollect(uid uint) (int64, error) {
	return client.ZCard(v.joinUserCollectKey(uid))
}

// GetCollect 分页获取用户收藏视频
func (v Video) GetCollect(uid uint, offset, number int64) ([]string, error) {
	return client.ZRevRange(v.joinUserCollectKey(uid), offset, offset+number-1)
}
