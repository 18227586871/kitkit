package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	Conf *viper.Viper
	one  sync.Once
)

func InitConfig() {
	one.Do(func() {
		v := viper.New()
		v.SetConfigName("conf")
		v.AddConfigPath("./config")
		v.SetConfigType("toml")
		if err := v.ReadInConfig(); err != nil {
			log.Println(err)
		}
		Conf = v
	})
}
