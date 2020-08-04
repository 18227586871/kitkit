package redis

import (
	"fmt"
	"testing"
)

// 这里写mock值
func TestName(t *testing.T) {
	redis := getRedis()
	fmt.Println(redis)
	redis.HSet()

}
