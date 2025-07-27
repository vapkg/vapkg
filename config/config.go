package config

import (
	"git
	"os"
	"vapkg/internal/core"
)

type vapkgConfig struct {
	RootPath    string `json:"VAPKG_ROOT,omitempty"`
	PackagePath string `json:"VAPKG_PATH,omitempty"`
	LogLevel    int    `json:"VAPKG_LOG_LEVEL,string,omitempty"`
	Env         string `json:"VAPKG_ENVIRONMENT,omitempty"`
}

var _ core.IConfig = (*Config)(nil)

type Config struct {
	container vapkgConfig
}

// NewConfig prod by default
func New() *Config {
	buf := vapkgConfig{}

	r := (enve.IEnveSource)(nil)
	if f, err := os.Open("./.env"); err == nil {
		defer f.Close()

		r = enve.NewReaderSource(f)
	}

	if err := enve.Parse(&buf, r, new(enve.EnvironSource)); err != nil {
		return nil
	}

	return &Config{buf}
}

func (c *Config) c() *vapkgConfig {
	return &c.container
}

func (c *Config) RootPath() string {
	return c.c().RootPath
}

func (c *Config) LogPath() string {
	return c.c().RootPath + "/logs"
}

func (c *Config) EnvType() core.EnvType {
	return core.EnvType(c.c().Env)
}

func (c *Config) LogLevel() core.LogType {
	return core.LogType(c.c().LogLevel)
}

func (c *Config) PackagePath() string {
	return c.c().PackagePath
}
