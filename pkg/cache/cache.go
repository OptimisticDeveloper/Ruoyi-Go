package cache

import (
	"ruoyi-go/pkg/cache/fileCache"
	"ruoyi-go/pkg/cache/redisCache"
	"ruoyi-go/pkg/cache/ristrettoCache"
	"time"
)

/**
c := cache.Cache("file")
	c.Put("113", "23", 1*time.Minute)
	result, _ := c.Get("113")
	fmt.Println(result)
*/

type CacheContract interface {
	Put(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
	Del(key string) (string, error)
	Clear() (string, error)
}

func Cache(driver string) CacheContract {
	switch driver {
	case "redis":
		return redisCache.NewRedisCache()
	case "file":
		return fileCache.NewFileCache()
	case "ristretto":
		return ristrettoCache.NewRistrettoCache()
	}
	return redisCache.NewRedisCache()
}
