package main

import (
	"log"
	"micro_service/config"
	"micro_service/internal/redis"
)

func main() {
	config.InitConfig()
	redis.InitRedis()
	log.Print()
}
