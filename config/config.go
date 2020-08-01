package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	Conf *viper.Viper
)

func InitConfig() {
	v := viper.New()
	v.SetConfigName("conf")
	v.AddConfigPath("./config")
	v.SetConfigType("toml")
	if err := v.ReadInConfig(); err != nil {
		log.Println(err)
	}
	Conf = v
}
