package dlock

// DLock 分布式锁抽象
type DLock interface {
	TryLock() error
	Release() (bool, error)
}
