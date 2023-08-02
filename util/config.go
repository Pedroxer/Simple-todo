package util

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBDriver      string
	DBAddress     string
	ServerAddress string
	TokenKey      string
	TokenDuration time.Duration
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("Config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("cannot load config: %s", err)
	}
	err = viper.Unmarshal(&config)
	return
}
