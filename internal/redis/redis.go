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

	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.GetString("redis.address"),
		Password: config.Conf.GetString("redis.password"), // no password set
		PoolSize: config.Conf.GetInt("redis.maxPoolSize"),
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	log.Println("Redis is Collection!!!", pong)
	redisCache = client
}

// 获取redis实例
func GetRedis() *redis.Client {
	return redisCache
}

//type redisStruct struct {
//	cache *redis.Client
//}

//func RCurd() *redisStruct {
//	return &redisStruct{
//		cache: getRedisPool(),
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
