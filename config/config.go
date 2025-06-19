package config

import (
	"os"
	"strconv"
	"vapkg/internal/core"
)

type Container struct {
	Bin      string       `mapstructure:"VAPKG_BIN"`
	LogLevel core.LogType `mapstructure:"VAPKG_LOGLEVEL"`
	Log      string       `mapstructure:"VAPKG_LOG"`
	Env      core.EnvType `mapstructure:"VAPKG_ENV"`
	Cache    string       `mapstructure:"VAPKG_CACHE"`
}

type Config struct {
	container Container
}

// NewConfig prod by default
func New(container Container) core.IConfig {
	return &Config{container: container}
}

func Get() core.IConfig {

	return New(Container{
		Bin:      GetBin(),
		Log:      GetLog(),
		LogLevel: GetLogType(),
		Env:      GetEnvironment(),
		Cache:    GetCache(),
	})
}

func (c *Config) BinFolder() string {
	return c.container.Bin
}

func (c *Config) LogFolder() string {
	return c.container.Log
}

func (c *Config) EnvType() core.EnvType {
	return c.container.Env
}

func (c *Config) LogLevel() core.LogType {
	return c.container.LogLevel
}

func (c *Config) CacheFolder() string {
	return c.container.Cache
}

func GetLogType() core.LogType {

	if os.Getenv("VAPKG_LOGLEVEL") == "" {
		return core.NoLog
	}

	if val, err := strconv.Atoi(os.Getenv("VAPKG_LOGLEVEL")); err == nil {
		return core.LogType(val)
	}

	return core.InfoLog
}

func GetLog() string {
	if val := os.Getenv("VAPKG_LOG"); val != "" {
		return val
	}

	return "log"
}

func GetBin() string {

	if val := os.Getenv("VAPKG_BIN"); val != "" {
		return val
	}

	return "bin"
}

func GetEnvironment() core.EnvType {
	if env := os.Getenv("VAPKG_ENV"); env == string(core.Development) {
		return core.EnvType(env)
	}

	return core.Production
}

func GetCache() string {
	if val := os.Getenv("VAPKG_CACHE"); val != "" {
		return val
	}

	return "bin/vapkg"
}
