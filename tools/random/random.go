package random

import (
	"math/rand"
	"time"
)

const (
	chars       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsLength = len(chars)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// New 生成指定字符数的随机字符串
func New(bits int) (randString string) {
	for i := 0; i < bits; i++ {
		randString += string(chars[rand.Intn(charsLength)])
	}
	return
}

// Default 生成16位随机字符串
func Default() string { return New(16) }
