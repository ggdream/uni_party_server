package rate

import (
	"errors"
	"time"
)

// leakyBucket 漏桶
type leakyBucket struct {
	size     int
	duration time.Duration

	waiter chan struct{}
	notify chan struct{}
}

func NewLeakyBucket(size, rps int) *leakyBucket {
	bucket := &leakyBucket{
		size:     size,
		duration: time.Second / time.Duration(rps),
		waiter:   make(chan struct{}, size),
		notify:   make(chan struct{}),
	}

	go bucket.Run()
	return bucket
}

func (b *leakyBucket) Run() {
	ticker := time.Tick(b.duration)
	for {
		<-ticker
		b.notify <- struct{}{}
	}
}

func (b *leakyBucket) Take() error {
	if len(b.waiter) == b.size {
		return errors.New("bucket is full")
	}

	b.waiter <- struct{}{}
	defer func() { <-b.waiter }()
	<-b.notify
	return nil
}
