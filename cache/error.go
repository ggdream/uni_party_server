package cache

import (
	"errors"

	"github.com/go-redis/redis/v8"
)

const Nil = redis.Nil

var (
	AddItAlreadyErr = errors.New("the user has been followed")
	DelItAlreadyErr = errors.New("the user has been unfollowed")
	MustGEZeroErr   = errors.New("values of offset and number must be ge zero")
	NotFoundErr     = errors.New("not found")
	MatchFailedErr  = errors.New("match failed")
)
