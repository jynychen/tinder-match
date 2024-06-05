package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServeAddr string `mapstructure:"SERVE_ADDR"`
	GinMode   string `mapstructure:"GIM_MODE"`
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("SERVE_ADDR", "127.0.0.1:8080")
	viper.SetDefault("GIM_MODE", "test")

	viper.AutomaticEnv()

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %s", err)
	}

	return &config, nil
}
