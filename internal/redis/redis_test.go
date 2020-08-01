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

	RCurd().Set("a", "dhy", 5*time.Minute)
	get := RCurd().Get("a")
	log.Println(get)
}
