package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	client *Redis
)

func Init() {
	client = New()
	if err := client.Ping(); err != nil {
		panic(err)
	}
}

// Redis 适用于本项目的Redis封装
type Redis struct {
	client  *redis.Client
	context context.Context
	timeout time.Duration
}

// New 实例化该封装
func New() *Redis {
	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	return &Redis{
		client:  redis.NewClient(options),
		context: context.Background(),
		timeout: 300 * time.Millisecond,
	}
}

// Ping 是否连接成功
func (r *Redis) Ping() error {
	ctx, cancelFunc := context.WithTimeout(r.context, 3*time.Second)
	defer cancelFunc()
	return r.client.Ping(ctx).Err()
}

// Exists 判断某个键是否存在
func (r *Redis) Exists(key string) (bool, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	res, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return res == 1, nil
}

// Del 删除一个或多个键
func (r *Redis) Del(key ...string) (int64, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.Del(ctx, key...).Result()
}

// Get get指令
func (r *Redis) Get(key string) *redis.StringCmd {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.Get(ctx, key)
}

// Set set指令
func (r *Redis) Set(key string, value interface{}, expire time.Duration) error {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.Set(ctx, key, value, expire).Err()
}

// LPush LPush指令
func (r *Redis) LPush(key string, value ...interface{}) error {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.LPush(ctx, key, value...).Err()
}

// LRem LRem指令
func (r *Redis) LRem(key string, count int64, value interface{}) error {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.LRem(ctx, key, count, value).Err()
}

// LLen LLen指令
func (r *Redis) LLen(key string) (int64, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.LLen(ctx, key).Result()
}

// LRange LRange指令
func (r *Redis) LRange(key string, start, end int64) ([]string, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.LRange(ctx, key, start, end).Result()
}

// HSet 给hash设置field和value
func (r *Redis) HSet(key, field string, value interface{}) error {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.HSet(ctx, key, field, value).Err()
}

// HMSet 给hash设置多组field和value
func (r *Redis) HMSet(key string, values ...interface{}) error {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.HSet(ctx, key, values...).Err()
}

// HGet 获取一组field和value
func (r *Redis) HGet(key, field string) *redis.StringCmd {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.HGet(ctx, key, field)
}

// HMGet 获取多组field和value
func (r *Redis) HMGet(key string, field ...string) ([]interface{}, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.HMGet(ctx, key, field...).Result()
}

// HGetAll 获取多组field和value
func (r *Redis) HGetAll(key string, field ...string) (map[string]string, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.HGetAll(ctx, key).Result()
}

// HDel 删除字段
func (r *Redis) HDel(key string, fields ...string) error {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.HDel(ctx, key, fields...).Err()
}

// HExists 判断hash的某个field是否存在
func (r *Redis) HExists(key, field string) (bool, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.HExists(ctx, key, field).Result()
}

// SAdd 给set添加成员
func (r *Redis) SAdd(key string, members ...interface{}) *redis.IntCmd {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.SAdd(ctx, key, members...)
}

// SRem 移除set里的成员
func (r *Redis) SRem(key string, members []interface{}) *redis.IntCmd {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.SRem(ctx, key, members...)
}

// SIsMember 判断该member是否为此set的一员
func (r *Redis) SIsMember(key string, members interface{}) (bool, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.SIsMember(ctx, key, members).Result()
}

// SMIsMember 判断该member是否为此set的一员
func (r *Redis) SMIsMember(key string, members ...interface{}) ([]bool, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.SMIsMember(ctx, key, members...).Result()
}

// SMembers 获取此set的所有元素
func (r *Redis) SMembers(key string) ([]string, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.SMembers(ctx, key).Result()
}

// ZAdd 给Sorted Sets添加成员
func (r *Redis) ZAdd(key string, members ...*redis.Z) error {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.ZAdd(ctx, key, members...).Err()
}

// ZAddNX 给Sorted Sets添加成员，不存在时才添加
func (r *Redis) ZAddNX(key string, members ...*redis.Z) (int64, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.ZAddNX(ctx, key, members...).Result()
}

// ZAddXX 给Sorted Sets添加成员，存在时才添加
func (r *Redis) ZAddXX(key string, members ...*redis.Z) (int64, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.ZAddXX(ctx, key, members...).Result()
}

// ZRem ZRem指令
func (r *Redis) ZRem(key string, members ...interface{}) (int64, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.ZRem(ctx, key, members...).Result()
}

// ZCard ZCard指令
func (r *Redis) ZCard(key string) (int64, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.ZCard(ctx, key).Result()
}

// ZRange ZRange指令
func (r *Redis) ZRange(key string, start, end int64) ([]string, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.ZRange(ctx, key, start, end).Result()
}

// ZRevRange ZRevRange指令
func (r *Redis) ZRevRange(key string, start, end int64) ([]string, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.ZRevRange(ctx, key, start, end).Result()
}

// ZRangeWithScores ZRangeWithScores指令
func (r *Redis) ZRangeWithScores(key string, start, end int64) ([]redis.Z, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.ZRangeWithScores(ctx, key, start, end).Result()
}

// ZRevRangeWithScores ZRevRangeWithScores指令
func (r *Redis) ZRevRangeWithScores(key string, start, end int64) ([]redis.Z, error) {
	ctx, cancelFunc := context.WithTimeout(r.context, r.timeout)
	defer cancelFunc()
	return r.client.ZRevRangeWithScores(ctx, key, start, end).Result()
}
