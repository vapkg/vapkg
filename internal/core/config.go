package core

type EnvType string

const (
	Production  EnvType = "prod"
	Development EnvType = "dev"
)

type IConfig interface {
	LogLevel() LogType
	LogPath() string
	RootPath() string
	EnvType() EnvType
	PackagePath() string
}
