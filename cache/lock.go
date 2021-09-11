package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"

	"gateway/tools/dlock"
)

const (
	DLockNameDelay = "dlock:delay"
)

var (
	DLockCannotRelease = errors.New("cannot release redis dlock")
)

var _ dlock.DLock = (*DistributedLock)(nil)
var redsyncClient *redSync

// distributedLock redis版分布式锁
type DistributedLock struct {
	mutex *redsync.Mutex

	timeout time.Duration
	context context.Context
}

// TryLock 尝试获取锁
func (l *DistributedLock) TryLock() error {
	return l.mutex.Lock()
}

// Release 释放锁
func (l *DistributedLock) Release() (bool, error) {
	return l.mutex.Unlock()
}

// TryLockContext 尝试获取锁
func (l *DistributedLock) TryLockContext() error {
	ctx, cancel := context.WithTimeout(l.context, l.timeout)
	defer cancel()

	return l.mutex.LockContext(ctx)
}

// UnlockContext 释放锁
func (l *DistributedLock) ReleaseContext() (bool, error) {
	ctx, cancel := context.WithTimeout(l.context, l.timeout)
	defer cancel()

	return l.mutex.UnlockContext(ctx)
}

// NewDLock 创建一个基于Redis的分布式锁实例
func NewDLock(name string) *DistributedLock {
	return &DistributedLock{
		mutex:   redsyncClient.newMutex(name),
		timeout: 600 * time.Millisecond,
		context: context.Background(),
	}
}

type redSync struct {
	rs *redsync.Redsync
}

// newMutex 申请某一分布式锁，以name区分
func (r *redSync) newMutex(name string) *redsync.Mutex {
	return r.rs.NewMutex(name)
}

func NewRedSync(client *redis.Client) *redSync {
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	return &redSync{rs: rs}
}
