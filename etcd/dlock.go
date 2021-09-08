package etcd

import (
	"context"
	"gateway/tools/dlock"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

var _ dlock.DLock = (*DLock)(nil)

// DLock etcd版分布式锁实现
// https://tangxusc.github.io/blog/2019/05/etcd-lock%E8%AF%A6%E8%A7%A3/
type DLock struct {
	pfx     string
	mutex   *concurrency.Mutex
	context context.Context
}

// NewDLock 实例化分布式锁
func NewDLock(pfx string) (*DLock, error) {
	return &DLock{
		pfx:     pfx,
		context: context.Background(),
	}, nil
}

// TryLock 获取锁
func (l *DLock) TryLock() error {
	ctx, cancel := context.WithTimeout(l.context, 3*time.Second)
	defer cancel()
	lease, err := client.client.Grant(ctx, 3)
	if err != nil {
		return err
	}

	session, err := concurrency.NewSession(client.client, concurrency.WithTTL(12), concurrency.WithLease(lease.ID))
	if err != nil {
		return err
	}

	l.mutex = concurrency.NewMutex(session, l.pfx)
	return l.mutex.Lock(ctx)
}

// Release 释放锁
func (l *DLock) Release() (bool, error) {
	ctx, cancel := context.WithTimeout(l.context, 1*time.Second)
	defer cancel()

	err := l.mutex.Unlock(ctx)
	return err == nil, err
}
