package redisutil

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	once     sync.Once
	instance *RedisClient
)

// RedisClient 封装 Redis 客户端
type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

// GetInstance 获取 Redis 客户端的单例
func GetInstance(addr string, password string, db int) *RedisClient {
	once.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})
		instance = &RedisClient{
			client: rdb,
			ctx:    context.Background(),
		}
	})
	return instance
}

// Set 存储键值对
func (r *RedisClient) Set(key string, value interface{}) error {
	return r.client.Set(r.ctx, key, value, 0).Err() // 0 表示永不过期
}

// Get 获取值
func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Close 关闭 Redis 客户端
func (r *RedisClient) Close() error {
	return r.client.Close()
}
