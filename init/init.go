package init

import (
	"micro_service/config"
	"micro_service/internal/mongodb"
	"micro_service/internal/redis"
)

func InitProject() {
	config.InitConfig()
	mongodb.InitMongo()
	redis.InitRedis()
}
