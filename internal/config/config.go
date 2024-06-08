package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

const path = "../../"

func ReadConfig(env string) (c Config) {
	fmt.Println(filepath.Dir(path))
	viper.AddConfigPath(filepath.Dir(path))
	viper.SetConfigName(env)
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
