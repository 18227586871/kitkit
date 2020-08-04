package redis

import (
	"github.com/go-redis/redis"
	"time"
)

//获取redis实例
func getRedis() *redis.Client {
	return redisCache
}

//
type redisStruct struct {
	cache *redis.Client
}

func RCurd() Iface {
	return &redisStruct{
		cache: getRedis(),
	}
}

// 先暂时订下来，后面可以直接用redis实例操作 set get
func (r *redisStruct) Set(key, val string, time time.Duration) {
	r.cache.Set(key, val, time)
}

func (r *redisStruct) Get(key string) string {
	return r.cache.Get(key).Val()
}
