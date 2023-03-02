package conf

import (
	"log"

	"github.com/spf13/viper"
)

func Init() {

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Fatal error config file: %v ", err.Error())
	}
}
