package base

import (
	"errors"
	"github.com/go-redis/redis"
	"pmo-test4.yz-intelligence.com/kit/cache/adapter"
	"time"
)

// Base 基础结构体，没实质功能
// 增加新接口，Base必须实现相应的接口
type Base struct{}

// Set 设置
func (*Base) Set(key string, value interface{}, expireation time.Duration) (bool, error) {
	return false, errors.New("未实现")
}

// Get 获取
func (*Base) Get(key string) (string, error) {
	return "", errors.New("未实现")
}

// Expire 设置过期时间
func (*Base) Expire(key string, expireation time.Duration) (bool, error) {
	return false, errors.New("未实现")
}

// Keys keys
func (*Base) Keys(pattern string) ([]string, error) {
	return nil, errors.New("未实现")
}

// Exists key是否存在
func (*Base) Exists(keys ...string) (int64, error) {
	return 0, errors.New("未实现")
}

// Delete 删除
func (*Base) Delete(keys ...string) (int64, error) {
	return 0, errors.New("未实现")
}

// HGet reids HGet
func (*Base) HGet(key string, field string) (string, error) {
	return "", errors.New("未实现")
}

// HSet redis HSet
func (*Base) HSet(key, field string, vals interface{}) (bool, error) {
	return false, errors.New("未实现")
}

// SetNX reids SetNX
func (*Base) SetNX(key string, value interface{}, expireation time.Duration) (bool, error) {
	return false, errors.New("未实现")
}

// ZAdd redis ZAdd
func (*Base) ZAdd(key string, members ...redis.Z) (int64, error) {
	return 0, errors.New("未实现")
}

// ZRem redis ZRem
func (*Base) ZRem(key string, members ...interface{}) (int64, error) {
	return 0, errors.New("未实现")
}

// ZRange reids ZRange
func (*Base) ZRange(key string, start, stop int64) ([]string, error) {
	return nil, errors.New("未实现")
}

// ZRangeWithScores ZRangeWithScores
func (*Base) ZRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return nil, errors.New("未实现")
}

// LPush LPush
func (*Base) LPush(key string, values ...interface{}) (bool, error) {
	return false, errors.New("未实现")
}

// GetAdapter 获取类型
func (*Base) GetAdapter() adapter.Adapter {
	return adapter.Unknow
}

// GetCore 获取操作对象实例,可配合GetAdapter进行断言
func (*Base) GetCore() (interface{}, error) {
	return nil, errors.New("未实现")
}

// Close 关闭连接
func (*Base) Close() error {
	return errors.New("未实现")
}
