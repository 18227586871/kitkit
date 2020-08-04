package redis

import "time"

type Iface interface {
	Set(key, val string, time time.Duration)
	Get(key string) string
}
