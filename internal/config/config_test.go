package config

import (
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"path/filepath"
	"testing"
)

const (
	Host     = "localhost"
	User     = "postgres"
	Password = "postgres"
	Port     = ":8080"
)

func TestReadConfig(t *testing.T) {
	viper.Reset()

	path := filepath.Join("..", "..")
	cfg := ReadConfig(path)
	assert.Equal(t, Host, cfg.DBHost, "host not equal")
	assert.Equal(t, User, cfg.DBUser, "user not equal")
	assert.Equal(t, Password, cfg.DBPassword, "password not equal")
	assert.Equal(t, Port, cfg.ServerPort, "port not equal")
}
