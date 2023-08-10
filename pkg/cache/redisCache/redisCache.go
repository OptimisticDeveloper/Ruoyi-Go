package redisCache

import (
	"context"
	"ruoyi-go/pkg/redis"
	"time"
)

type redisCache struct {
}

func NewRedisCache() *redisCache {
	return &redisCache{}
}

func (r redisCache) Put(key string, value string, ttl time.Duration) error {
	_, err := redis.Client().Set(context.TODO(), key, value, ttl).Result()
	return err
}

func (r redisCache) Get(key string) (string, error) {
	return redis.Client().Get(context.TODO(), key).Result()
}

func (r redisCache) Del(key string) (string, error) {
	_, error := redis.Client().Del(context.TODO(), key).Result()
	return "", error
}

func (r redisCache) Clear() (string, error) {
	return "", nil
}
