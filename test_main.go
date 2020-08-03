package main

import (
	"fmt"
	"micro_service/internal/redis"
)

func main() {
	redis.InitRedis()
	getRedis := redis.GetRedis().LPush("a", "aaa").Val()

	fmt.Println(getRedis)
}
