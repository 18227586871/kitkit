package redis

import (
	"log"
	"micro_service/config"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	config.InitConfig()
	InitRedis()

	GetRedis().Set("a", "aaa", time.Second)
	get := GetRedis().Get("a").Val()
	log.Println(get)
}
