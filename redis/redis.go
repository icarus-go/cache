package redis

import (
	"fmt"
	"time"

	"amiba.fun/amiba/go-cache/adapter"
	"amiba.fun/amiba/go-cache/base"
	"amiba.fun/amiba/go-cache/config"
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

// Keys keys
func (c *Client) Keys(pattern string) ([]string, error) {
	stringSliceCMD := c.client.Keys(pattern)
	return stringSliceCMD.Result()
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
