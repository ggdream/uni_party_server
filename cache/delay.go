package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"time"
)

// DelayQueue 延时队列
type DelayQueue struct{}

// joinQueueKey 拼接延时队列的key：ZSet
func (d DelayQueue) joinQueueKey() string {
	return "delay:queue"
}

// Add 添加延时任务
func (d DelayQueue) Add(eid string, time int64) error {
	el := redis.Z{
		Score:  float64(time),
		Member: eid,
	}

	return client.ZAdd(d.joinQueueKey(), &el)
}

// Poll 轮询判断
func (d DelayQueue) Poll(callback func([]string) error, errHandler func(error)) {
	tick := time.Tick(time.Second * 1)

	for {
		err := d.handle(callback)
		if err != nil {
			errHandler(err)
		}

		<-tick
	}
}

// handle 用分布式锁获取、处理、删除消息
func (d DelayQueue) handle(callback func([]string) error) (err error) {
	dlock := NewDLock(DLockNameDelay)
	err = dlock.TryLock()
	if err != nil {
		return
	}
	defer func(dlock *DistributedLock) {
		var ok bool
		ok, err = dlock.Release()
		if err != nil && !ok {
			err = DLockCannotRelease
		}
	}(dlock)

	value, err := client.ZRange(d.joinQueueKey(), 0, time.Now().Unix())
	if err != nil {
		return
	}

	err = callback(value)
	if err != nil {
		return
	}

	valueInterface := make([]interface{}, len(value))
	for i, v := range value {
		valueInterface[i] = v
	}
	count, err := client.ZRem(d.joinQueueKey(), valueInterface...)

	if int(count) < len(value) {
		err = errors.New("don't delete them completely")
	}

	return
}
