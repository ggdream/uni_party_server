package cache

import (
	"fmt"
	"gateway/tools/hashids"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	eventHash = hashids.New("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", "abcd", 8)
)

// Event 事务消息
// 消息统计：Get使用 Sorted Sets；Attend使用 Sorted Sets；Watch使用 Hash
// 用户统计：Get使用 Sorted Sets；Attend使用 Sorted Sets
// 用户发布使用：List；用户订阅使用：List
type Event struct{}

// joinCountGetKey 获取Get消息的集合元素数量的string-key
func (e Event) joinCountGetKey(eid string) string {
	return fmt.Sprintf("event:count:get:%s", eid)
}

// joinCountAttendKey 获取关注消息的列表元素数量的string-key
func (e Event) joinCountAttendKey(eid string) string {
	return fmt.Sprintf("event:count:attend:%s", eid)
}

// joinCountWatchKey 获取查看消息计数的哈希的string-key
func (e Event) joinCountWatchKey(eid string) string {
	eventIntID, _ := eventHash.Decode(eid)
	return fmt.Sprintf("event:count:watch:%d", eventIntID / 4096)
}

// joinUserGetKey 获取用户Get的消息集合的string-key
func (e Event) joinUserGetKey(uid uint) string {
	return fmt.Sprintf("event:user:get:%d", uid)
}

// joinUserAttendKey 获取用户关注消息集合的string-key
func (e Event) joinUserAttendKey(uid uint) string {
	return fmt.Sprintf("event:user:attend:%d", uid)
}

// joinUserPubKey 获取用户发布消息集合的string-key
func (e Event) joinUserPubKey(uid uint) string {
	return fmt.Sprintf("event:user:pub:%d", uid)
}

// joinUserSubKey 获取用户订阅消息集合的string-key
func (e Event) joinUserSubKey(uid uint) string {
	return fmt.Sprintf("event:user:sub:%d", uid)
}


// EventCountGet 获取消息的Get数量
func (e Event) EventCountGet(eid string) (int64, error) {
	return client.ZCard(e.joinCountGetKey(eid))
}

// EventCountAttend 获取消息的关注数量
func (e Event) EventCountAttend(eid string) (int64, error) {
	return client.LLen(e.joinCountAttendKey(eid))
}

// EventCountWatch 获取消息的查看数量
func (e Event) EventCountWatch(eid string) (int64, error) {
	return client.HGet(e.joinCountWatchKey(eid), eid).Int64()
}

// UserCountGet 获取用户消息的Get数量
func (e Event) UserCountGet(uid uint) (int64, error) {
	return client.ZCard(e.joinUserGetKey(uid))
}

// UserCountAttend 获取消息的关注数量
func (e Event) UserCountAttend(uid uint) (int64, error) {
	return client.ZCard(e.joinUserAttendKey(uid))
}

// UserCountPub 获取用户消息的Get数量
func (e Event) UserCountPub(uid uint) (int64, error) {
	return client.LLen(e.joinUserPubKey(uid))
}

// Publish 发布消息
// publisher 为发布者uid；subscribers为订阅者uid
func (e Event) Publish(publisher uint, eid string, subscribers ...uint) error {
	for _, uid := range subscribers {
		if err := client.LPush(e.joinUserSubKey(uid), eid); err != nil {
			return err
		}
	}
	return client.LPush(e.joinUserPubKey(publisher), eid)
}

// Delete 删除消息
// publisher 为发布者uid；subscribers为订阅者uid
func (e Event) Delete(publisher uint, eid string, subscribers ...uint) error {
	for _, uid := range subscribers {
		if err := client.LRem(e.joinUserSubKey(uid), 0, eid); err != nil {
			return err
		}
	}
	return client.LRem(e.joinUserPubKey(publisher), 0, eid)
}

// QueryPub 分页查询发布消息
func (e Event) QueryPub(uid uint, offset, number int64) ([]string, error) {
	return client.LRange(e.joinUserPubKey(uid), offset, number - 1)
}

// QuerySub 分页查询订阅消息
func (e Event) QuerySub(uid uint, offset, number int64) ([]string, error) {
	return client.LRange(e.joinUserSubKey(uid), offset, number - 1)
}


// DoGet 用户Get消息
func (e Event) DoGet(uid uint, eid string) error {
	el := &redis.Z{
		Score:  float64(time.Now().UnixNano()),
		Member: uid,
	}
	hadAdded, err := client.ZAddNX(e.joinCountGetKey(eid), el)
	if err != nil {
		return err
	}
	// 如果之前已经添加过，则返回错误 FollowItAlreadyErr
	if hadAdded == 0 {
		return AddItAlreadyErr
	}

	el.Member = eid	// 为了复用结构体实例
	return client.ZAdd(e.joinUserGetKey(uid), el)
}

// DoAttend 用户关注消息
func (e Event) DoAttend(uid uint, eid string) error {
	el := &redis.Z{
		Score:  float64(time.Now().UnixNano()),
		Member: uid,
	}
	hadAdded, err := client.ZAddNX(e.joinCountAttendKey(eid), el)
	if err != nil {
		return err
	}
	// 如果之前已经添加过，则返回错误 FollowItAlreadyErr
	if hadAdded == 0 {
		return AddItAlreadyErr
	}

	el.Member = eid	// 为了复用结构体实例
	return client.ZAdd(e.joinUserAttendKey(uid), el)
}

// DoUnAttend 用户取关消息
func (e Event) DoUnAttend(uid uint, eid string) error {
	if err := client.ZRem(e.joinCountAttendKey(eid), uid); err != nil {
		return err
	}
	return client.ZRem(e.joinUserAttendKey(uid), eid)
}

// GetAttend 获取用户关注的消息
func (e Event) GetAttend(uid uint, offset, number int64) ([]string, error) {
	if offset < 0 || number < 0 {
		return nil, MustGEZeroErr
	}
	if number != 0 {
		number--
	}
	return client.ZRevRange(e.joinUserAttendKey(uid), offset, number)
}

// GetPublish 获取用户发布的消息
func (e Event) GetPublish(uid uint, offset, number int64) ([]string, error) {
	if offset < 0 || number < 0 {
		return nil, MustGEZeroErr
	}
	if number != 0 {
		number--
	}
	return client.ZRevRange(e.joinUserPubKey(uid), offset, number)
}
