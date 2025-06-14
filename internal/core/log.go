package core

import (
	"io"
	"time"
)

type LogType uint8

const (
	NoLog    LogType = 0x0
	ErrLog   LogType = 0x1
	WarnLog  LogType = 0x2
	InfoLog  LogType = 0x4
	DebugLog LogType = 0x8
)

type ILogger interface {
	Type() LogType
	Writer() io.Writer
	Close()
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
}

func GetActualLogFile() string {
	return "vapkg_" + time.Now().Format("2006-01-02") + ".txt"
}
