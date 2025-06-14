package core

import (
	"github.com/spf13/viper"
	"vapkg/internal/utils"
)

const source = "./.env"

type EnvType string

const (
	Production  EnvType = "prod"
	Development EnvType = "dev"
)

type Config struct {
	BinFolder string        `mapstructure:"VAPKG_BIN"`
	LogLevel  utils.LogType `mapstructure:"VAPKG_LOGLEVEL"`
	LogFolder string        `mapstructure:"VAPKG_LOG"`
	EnvType   EnvType       `mapstructure:"VAPKG_ENV"`
}

// NewConfig prod by default
func NewConfig() *Config {
	return &Config{
		BinFolder: "bin",
		LogLevel:  utils.InfoLog,
		LogFolder: "/var/log",
		EnvType:   Production,
	}
}

func GetConfig() (*Config, error) {
	config := NewConfig()

	viper.SetConfigFile(source)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
