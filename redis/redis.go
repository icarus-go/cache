package redis

import (
	"fmt"
	"github.com/icarus-go/cache/adapter"
	"github.com/icarus-go/cache/base"
	"github.com/icarus-go/cache/config"
	"time"

	"github.com/go-redis/redis"
)

// Client 结构体
type Client struct {
	base.Base
	client *redis.Client
	config *config.Config
}

// New 实例化
func New(cfg *config.Config) (*Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Host + ":" + cfg.Port,
		Password:     cfg.Password,
		DB:           0,
		MinIdleConns: 5,
		//ReadTimeout: time.Minute,
	})

	if err := client.Ping().Err(); err != nil {
		panic(fmt.Sprintf("redis 连接失败 [%v]", err))
	}
	rc := &Client{}
	rc.client = client
	rc.config = cfg
	return rc, nil
}

//Get get
func (c *Client) Get(key string) (string, error) {
	return c.client.Get(key).Result()
}

// Set 设置 成功返回true 否则返回false
func (c *Client) Set(key string, value interface{}, expireation time.Duration) (bool, error) {

	ret := c.client.Set(key, value, expireation)

	err := ret.Err()
	if err != nil {
		return false, err
	}

	if ret.Val() == "OK" {
		return true, nil
	}

	return false, nil
}

// Delete 删除
func (c *Client) Delete(keys ...string) (int64, error) {
	intCMD := c.client.Del(keys...)
	return intCMD.Result()
}

// Keys keys 线上环境，禁止使用Keys
func (c *Client) Keys(pattern string) ([]string, error) {
	stringSliceCMD := c.client.Keys(pattern)
	return stringSliceCMD.Result()
}

func (c *Client) Scan(cursor uint64, match string, count int64) ([]string, error) {
	var value []string
	result := c.client.Scan(cursor, match, count)
	if result.Err() != nil {
		return nil, result.Err()
	}
	value, cursor = result.Val()
	return value, nil
}

// SetNX SetNX
func (c *Client) SetNX(key string, value interface{}, expireation time.Duration) (bool, error) {
	boolCMD := c.client.SetNX(key, value, expireation)
	return boolCMD.Result()
}

// Close 关闭连接
func (c *Client) Close() error {
	return c.client.Close()
}

// Expire 设置过期时间
func (c *Client) Expire(key string, expireation time.Duration) (bool, error) {
	boolCMD := c.client.Expire(key, expireation)

	err := boolCMD.Err()
	if err != nil {
		return false, err
	}

	return boolCMD.Val(), nil
}

// Exists 是否存在
func (c *Client) Exists(keys ...string) (int64, error) {
	intCMD := c.client.Exists(keys...)
	err := intCMD.Err()
	if err != nil {
		return 0, err
	}

	return intCMD.Val(), nil
}

// GetCore 返回redis实例
func (c *Client) GetCore() (interface{}, error) {
	return c.client, nil
}

// GetAdapter 返回类型
func (c *Client) GetAdapter() adapter.Adapter {
	return c.config.Adapter
}
