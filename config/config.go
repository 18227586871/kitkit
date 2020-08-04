package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Keys  string
	Mongo mongoConf
	Http  httpConf
	Mysql mysqlConf
	Redis redisConf
}

// Mongo
type mongoConf struct {
	Address     string
	MaxPoolSize uint64
}

// Http
type httpConf struct {
	Listen string
}

// Mysql
type mysqlConf struct {
	Address     string
	MaxOpenConn int
	MaxIdleConn int
}

// Redis
type redisConf struct {
	Address     string
	MaxPoolSize int
	Password    string
}

var conf Config

func init() {
	v := viper.New()

	v.SetConfigName("conf")     // 配置文件的名字
	v.SetConfigType("toml")     // 配置文件的类型
	v.AddConfigPath("./config") // 配置文件的路径

	if err := v.ReadInConfig(); err != nil {
		panic(err)
		return
	}

	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}
}

func GetConf() Config {
	return conf
}
