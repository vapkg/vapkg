package utils

import (
	"fmt"
	"io"
	"os"
	"path"
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

func GetActualLogFile() string {
	return "vapkg_" + time.Now().Format("2000-01-02") + ".txt"
}

type ILogger interface {
	Type() LogType
	Writer() io.Writer
	Close()
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
}

type Logger struct {
	writer  *os.File
	logType LogType
}

func CreateLogger(w *os.File, logType LogType) Logger {
	return Logger{writer: w, logType: logType}
}

func NewLogger(w *os.File, logType LogType) ILogger {
	return &Logger{writer: w, logType: logType}
}

func CreateActualLogger(dir string, t LogType) (ILogger, error) {
	var err error
	var file *os.File
	if file, err = os.OpenFile(path.Join(dir, GetActualLogFile()), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666); err != nil {
		return nil, err
	}

	return NewLogger(file, t), nil
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

func (l *Logger) Writer() io.Writer {
	return l.writer
}

func (l *Logger) Close() {
	if l.writer != nil {
		if err := l.writer.Close(); err != nil {
			return
		}
	}
}
