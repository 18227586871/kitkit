package redis

import (
	"github.com/go-redis/redis"
	"log"
	"micro_service/config"
	"sync"
)

var (
	RedisClient *redis.Client
	once        sync.Once
)

func InitRedis() {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     config.Conf.GetString("redis.address"),
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		pong, err := client.Ping().Result()
		if err != nil {
			panic(err)
		}
		log.Println("Redis is Collection!!!", pong)
		RedisClient = client
	})
}
