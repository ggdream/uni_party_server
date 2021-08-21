package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// Follow 关注与粉丝
// 关注列表使用 Sorted Sets
// 粉丝列表使用 List
// value为uid，score为操作时间戳
// 交并差得到关注关系，分段获取实现分页
type Follow struct{}

// joinUserFollowingKey 用户关注的string-key
func (f Follow) joinUserFollowingKey(uid uint) string {
	return fmt.Sprintf("follow:ing:%d", uid)
}

// joinUserFollowersKey 用户粉丝的string-key
func (f Follow) joinUserFollowersKey(uid uint) string {
	return fmt.Sprintf("follow:ers:%d", uid)
}

// CountFollowing 获取用户的关注数量
func (f Follow) CountFollowing(uid uint) (int64, error) {
	return client.ZCard(f.joinUserFollowingKey(uid))
}

// CountFollowers 获取用户的粉丝数量
func (f Follow) CountFollowers(uid uint) (int64, error) {
	return client.LLen(f.joinUserFollowersKey(uid))
}

// GetFollowing 获取用户的关注者信息
func (f Follow) GetFollowing(uid uint, offset, number int64) ([]redis.Z, error) {
	return client.ZRevRangeWithScores(f.joinUserFollowingKey(uid), offset, offset+number-1)
}

// GetFollowers 获取用户的粉丝信息
func (f Follow) GetFollowers(uid uint, offset, number int64) ([]string, error) {
	return client.LRange(f.joinUserFollowersKey(uid), offset, offset+number-1)
}

// Follow 关注他人
func (f Follow) Follow(fromUID, toUID uint) error {
	el := &redis.Z{
		Score:  float64(time.Now().UnixNano()),
		Member: toUID,
	}
	// 添加到自己的关注集合
	hadFollowed, err := client.ZAddNX(f.joinUserFollowingKey(fromUID), el)
	if err != nil {
		return err
	}
	// 如果之前已经关注过，则返回错误 FollowItAlreadyErr
	if hadFollowed == 0 {
		return AddItAlreadyErr
	}

	return client.LPush(f.joinUserFollowersKey(toUID), fromUID)
}

// UnFollow 取关他人
func (f Follow) UnFollow(fromUID, toUID uint) error {
	// 从自己的关注集合移除对方
	v, err := client.ZRem(f.joinUserFollowingKey(fromUID), toUID)
	if err != nil {
		return err
	}
	if v == 0 {
		return DelItAlreadyErr
	}

	// 从对方的粉丝列表移除自己
	return client.LRem(f.joinUserFollowersKey(toUID), 0, fromUID)
}
