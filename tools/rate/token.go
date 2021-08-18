package rate

import (
	"errors"
	"time"
)

// tokenBucket 令牌桶
type tokenBucket struct {
	size     int
	duration time.Duration

	bucket chan struct{}
}

func NewTokenBucket(size, rps int) *tokenBucket {
	bucket := &tokenBucket{
		size:     size,
		duration: time.Second / time.Duration(rps),
		bucket:   make(chan struct{}, size),
	}

	go bucket.Run()
	return bucket
}

func (b *tokenBucket) Run() {
	ticker := time.Tick(b.duration)
	for {
		<-ticker
		b.bucket <- struct{}{}
	}
}

func (b *tokenBucket) Take() error {
	if len(b.bucket) == 0 {
		return errors.New("bucket is empty")
	}

	<-b.bucket
	return nil
}
