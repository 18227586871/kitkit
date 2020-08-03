package redis

import (
	"github.com/go-redis/redis"
	"log"
	"micro_service/config"
)

var (
	redisCache *redis.Client
)

func InitRedis() {

	redisOptions := &redis.Options{
		ReadTimeout: -1,
	}
	redisOptions.Addr = config.GetConf().Redis.Address
	if config.GetConf().Redis.Password != "" {
		redisOptions.Password = config.GetConf().Redis.Password
	}
	redisOptions.DB = 0
	if config.GetConf().Redis.MaxPoolSize > 0 {
		redisOptions.PoolSize = config.GetConf().Redis.MaxPoolSize
	}
	redisCache = redis.NewClient(redisOptions)

	pong, err := redisCache.Ping().Result()
	if err != nil {
		panic(err)
	}

	log.Println("Redis is Collection!!!", pong)
}

//获取redis实例
func GetRedis() *redis.Client {
	return redisCache
}

//
//type redisStruct struct {
//	cache *redis.Client
//}
//
//func RCurd() *redisStruct {
//	return &redisStruct{
//		cache: GetRedis(),
//	}
//}
//
//// 先暂时订下来，后面可以直接用redis实例操作 set get
//func (r *redisStruct) Set(key, val string, time time.Duration) {
//	r.cache.Set(key, val, time)
//}
//
//func (r *redisStruct) Get(key string) string {
//	return r.cache.Get(key).Val()
//}
