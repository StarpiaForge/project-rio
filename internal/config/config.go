// Package config contains structures and methods of application configuration.
package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

var (
	ErrReadConfig      = errors.New("cannot read config")
	ErrUnmarshalConfig = errors.New("cannot unmarshal config")
)

type Config struct {
	APIKey string
}

func Load(configPath string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	viper.SetEnvPrefix("RIO")
	viper.BindEnv("api_key", "RIO_API_KEY")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrReadConfig, err)
	}

	c := new(Config)
	if err := viper.Unmarshal(c); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnmarshalConfig, err)
	}

	return c, nil
}
