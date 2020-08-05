package main

import (
	"micro_service/config"
	"micro_service/internal/pkg/redis"
	"time"
)

func main() {
	var c config.Config
	config.Conf().InitConfig("conf", "toml", "./config", &c)
	redis.InitRedis(c.Redis.Address, c.Redis.Password, c.Redis.MaxPoolSize)
	redis.RCurd().Set("a", "aaaaaaaa", 8*time.Minute)
}
