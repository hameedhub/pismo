package config

import (
	"github.com/spf13/viper"
	"log"
)

func ReadConfig(path string) (c Config) {
	viper.AddConfigPath(path)
	viper.AddConfigPath("/app/")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("cannot read read config error: ", err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal("cannot unmarshal file error: ", err)
	}
	return c
}
