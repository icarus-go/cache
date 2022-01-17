package cache

import (
	"errors"
	"github.com/icarus-go/cache/adapter"
	"github.com/icarus-go/cache/config"
	"github.com/icarus-go/cache/redis"
	"time"

	rediscore "github.com/go-redis/redis"
)

// ICache 缓存接口
type ICache interface {
	// Set 设置
	Set(key string, value interface{}, expireation time.Duration) (bool, error)

	// Get 获取
	Get(key string) (string, error)

	// Exists key 是否存在
	Exists(keys ...string) (int64, error)

	// Delete 删除
	Delete(keys ...string) (int64, error)

	// Expire 设置过期时间
	Expire(key string, expireation time.Duration) (bool, error)

	// Keys keys
	Keys(pattern string) ([]string, error)

	// HGet reids HGet
	HGet(key string, field string) (string, error)

	// HSet redis HSet
	HSet(key, field string, vals interface{}) (bool, error)

	// SetNX reids SetNX
	SetNX(key string, value interface{}, expireation time.Duration) (bool, error)

	// ZAdd redis ZAdd
	ZAdd(key string, members ...rediscore.Z) (int64, error)

	// ZRem redis ZRem
	ZRem(key string, members ...interface{}) (int64, error)

	// ZRange reids ZRange
	ZRange(key string, start, stop int64) ([]string, error)

	// ZRangeWithScores ZRangeWithScores
	ZRangeWithScores(key string, start, stop int64) ([]rediscore.Z, error)

	// LPusth redis LPush
	LPush(key string, values ...interface{}) (bool, error)

	// GetAdapter 获取类型
	GetAdapter() adapter.Adapter

	// GetCore 获取操作对象实例,可配合GetAdapter进行断言
	GetCore() (interface{}, error)

	// Scan 定量获取Key
	Scan(cursor uint64, match string, count int64) ([]string, error)

	// Close 关闭连接
	Close() error
}

// New 实例化
func New(cfg *config.Config) (ICache, error) {
	switch cfg.Adapter {
	case adapter.Redis:
		client, err := redis.New(cfg)
		if err != nil {
			return nil, err
		}
		return client, nil
	default:
		return nil, errors.New("不存在的缓存适配器")
	}
}
