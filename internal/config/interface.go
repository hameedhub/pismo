package config

type Config struct {
	DBHost             string `mapstructure:"DB_HOST"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBName             string `mapstructure:"DB_NAME"`
	DBPort             string `mapstructure:"DB_PORT"`
	ServerPort         string `mapstructure:"SERVER_PORT"`
	ServerIdleTimeout  int32  `mapstructure:"SERVER_IDLE_TIMEOUT"`
	ServerReadTimeout  int32  `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout int32  `mapstructure:"SERVER_WRITE_TIMEOUT"`
}
