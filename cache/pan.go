package cache

import (
	"context"
	"fmt"
	"time"

	"gateway/tools/hashids"
	"gateway/tools/random"

	"github.com/go-redis/redis/v8"
)

const (
	PanShareHour = int8(iota)
	PanShareDay
	PanShareDay3
	PanShareWeek
	PanShareMonth
	PanSharePermanent
)

var panHashIDS = hashids.New("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "Y3J4X0GhwRCm00Qr", 8)

type Pan struct{}

func (p Pan) joinRecordShareKey(rand, password string) string {
	return fmt.Sprintf("pan:record:share:%s:%s", rand, password)
}

// joinxRecordShareKey string: 以joinUserShareKey的value为参数
func (p Pan) joinxRecordShareKey(randPwd string) string {
	return fmt.Sprintf("pan:record:share:%s", randPwd)
}

// joinTokenKey string:获取文件内容所必须的token
func (p Pan) joinTokenKey(token string) string {
	return fmt.Sprintf("pan:record:token:%s", token)
}

// joinUserShareKey hash：获取某个用户分享的所有文件，在下一次获取分享时
// field为文件/文件夹hash，value为`rand:password`
func (p Pan) joinUserShareKey(uid uint) string {
	return fmt.Sprintf("pan:user:share:%d", uid)
}

func (p Pan) transformExpireTime(expireTime int8) time.Duration {
	switch expireTime {
	case PanShareHour:
		return time.Hour
	case PanShareDay:
		return 24 * time.Hour
	case PanShareDay3:
		return 3 * 24 * time.Hour
	case PanShareWeek:
		return 7 * 24 * time.Hour
	case PanShareMonth:
		return 30 * 24 * time.Hour
	case PanSharePermanent:
		return 0
	default:
		return time.Hour
	}
}

// Share 分享文件/文件夹
func (p Pan) Share(uid uint, hash string, expireTime int8) (string, string, error) {
	rand := panHashIDS.Encode(uint(time.Now().Unix()*100000000) + uid)
	password := random.New(4)
	postTime := p.transformExpireTime(expireTime)

	err := client.Set(p.joinRecordShareKey(rand, password), hash, postTime)
	if err != nil {
		return "", "", err
	}
	err = client.HSet(p.joinUserShareKey(uid), hash, fmt.Sprintf("%s:%s", rand, password))
	if err != nil {
		return "", "", err
	}

	return rand, password, nil
}

// Unshare 取消分享
func (p Pan) Unshare(uid uint, hash string) error {
	randPwd, err := client.HGet(p.joinUserShareKey(uid), hash).Result()
	if err != nil {
		return err
	}
	err = client.HDel(p.joinUserShareKey(uid), hash)
	if err != nil {
		return err
	}

	_, err = client.Del(p.joinxRecordShareKey(randPwd))
	return err
}

// GetShareArchive 获取用户自己的分享hash值
func (p Pan) GetShareArchive(uid uint) ([]string, error) {
	var aliveList, expireList []string
	key := p.joinUserShareKey(uid)

	transaction := func(tx *redis.Tx) error {
		pipe := tx.TxPipeline()
		defer pipe.Close()

		shareMap, err := pipe.HGetAll(context.Background(), key).Result()
		if err != nil {
			return err
		}

		for k, v := range shareMap {
			exists, err := pipe.Exists(context.Background(), p.joinxRecordShareKey(v)).Result()
			if err != nil {
				return err
			}

			if exists == 0 {
				expireList = append(expireList, k)
			} else {
				aliveList = append(aliveList, k)
			}
		}

		err = pipe.HDel(context.Background(), key, expireList...).Err()
		if err != nil {
			return err
		}

		_, err = pipe.Exec(context.Background())
		if err != nil {
			return pipe.Discard()
		}
		return nil
	}

	err := client.client.Watch(client.context, transaction, key)
	if err != nil {
		return nil, err
	}

	return aliveList, nil
}

// GetShare 接收者输入密码，查看文件
func (p Pan) GetShare(rand, password string) (string, error) {
	return client.Get(p.joinRecordShareKey(rand, password)).Result()
}

// SetToken 允许下载文件的令牌
func (p Pan) SetToken(token, rand string) error {
	return client.Set(p.joinTokenKey(token), rand, 3*time.Hour)
}

// CheckToken 检测令牌是否存在
func (p Pan) CheckToken(token, rand string) (bool, error) {
	r, err := client.Get(p.joinTokenKey(token)).Result()
	if err != nil {
		return false, err
	}

	return r == rand, nil
}
