package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Session string
type Session struct{}

func (s Session) joinKey(uid uint) string {
	return fmt.Sprintf("user:session:%d", uid)
}

// Set 按uid存储session
func (s Session) Set(uid uint, session string) error {
	return client.Set(s.joinKey(uid), session, 7*24*time.Hour)
}

// Verify 鉴别uid和session的对应关系
func (s Session) Verify(uid uint, session string) (bool, error) {
	value, err := client.Get(s.joinKey(uid)).Result()
	if err != nil {
		if err == redis.Nil {
			return false, NotFoundErr
		} else {
			return false, err
		}
	}

	if value != session {
		return false, MatchFailedErr
	}
	return true, nil
}
