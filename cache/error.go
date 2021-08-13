package cache

import (
	"errors"
	"github.com/go-redis/redis/v8"
)

const Nil = redis.Nil

var (
	AddItAlreadyErr = errors.New("the user has been followed")
	MustGEZeroErr = errors.New("values of offset and number must be ge zero")
)
