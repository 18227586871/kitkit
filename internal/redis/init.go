package redis

import (
	"github.com/go-redis/redis"
	"log"
	"micro_service/config"
)

var (
	redisClient *redis.Client
)

func InitRedis() {

	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.GetString("redis.address"),
		Password: "", // no password set
		DB:       0,  // use default DB
		//PoolSize:     config.Conf.GetInt("redis.maxPoolSize"),
		PoolSize:     100,
		MinIdleConns: 15,
	})

	//defer client.Close()
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	log.Println("Redis is Collection!!!", pong)
	redisClient = client
}
