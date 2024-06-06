package config

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"testing"
)

const (
	Host     = "localhost"
	User     = "postgres"
	Password = "postgres"
	Port     = "5000"
)

func TestConfigDefaults(t *testing.T) {
	viper.Reset()

	cfg := ReadConfig("dev")
	fmt.Println(cfg.DBHost)
	assert.Equal(t, Host, cfg.DBHost, "host not equal")
	assert.Equal(t, User, cfg.DBUser, "user not equal")
	assert.Equal(t, Password, cfg.DBPassword, "password not equal")
	assert.Equal(t, Port, cfg.ServerPort, "port not equal")
}
