package main

import (
	"fmt"
	"micro_service/internal/redis"
)

func main() {
	redis.InitRedis()

	get := redis.RCurd().Get("ddddd")
	fmt.Println(get)
}
