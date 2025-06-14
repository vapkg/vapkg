package core

type EnvType string

const (
	Production  EnvType = "prod"
	Development EnvType = "dev"
)

type IConfig interface {
	LogLevel() LogType
	LogFolder() string
	BinFolder() string
	EnvType() EnvType
}
