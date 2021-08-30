// Package redis sorted set 有序列集合
package redis

import "github.com/go-redis/redis"

// ZAdd ZAdd
func (c *Client) ZAdd(key string, members ...redis.Z) (int64, error) {
	intCMD := c.client.ZAdd(key, members...)
	return intCMD.Result()
}

// ZRem ZRem
func (c *Client) ZRem(key string, members ...interface{}) (int64, error) {
	intCMD := c.client.ZRem(key, members...)
	return intCMD.Result()
}

// ZRange ZRange
func (c *Client) ZRange(key string, start, stop int64) ([]string, error) {
	stringSliceCMD := c.client.ZRange(key, start, stop)
	return stringSliceCMD.Result()
}

// ZRangeWithScores ZRangeWithScores
func (c *Client) ZRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	zSliceCMD := c.client.ZRangeWithScores(key, start, stop)
	return zSliceCMD.Result()
}
