package utils

import (
	"fmt"
	"os"
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
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type Logger struct {
	logType LogType
}

func CreateLogger(logType LogType) Logger {
	return Logger{logType}
}

func NewLogger(logType LogType) ILogger {
	return &Logger{logType: logType}
}

func (l *Logger) Type() LogType {
	return l.logType
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if (l.logType & DebugLog) != 0 {
		fmt.Fprintf(os.Stdout, "[DEBUG] "+format+"\n", args...)
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if (l.logType & InfoLog) != 0 {
		fmt.Fprintf(os.Stdout, "[INFO] "+format+"\n", args...)
	}
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if (l.logType & WarnLog) != 0 {
		fmt.Fprintf(os.Stdout, "[WARN] "+format+"\n", args...)
	}
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if (l.logType & ErrLog) != 0 {
		fmt.Fprintf(os.Stdout, "[ERROR] "+format+"\n", args...)
	}
}
