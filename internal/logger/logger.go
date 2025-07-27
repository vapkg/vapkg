package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"
	"vapkg/internal/core"
)

type Logger struct {
	writer  *os.File
	logType core.LogType
}

func Create(w *os.File, logType core.LogType) Logger {
	return Logger{writer: w, logType: logType}
}

func New(w *os.File, logType core.LogType) *Logger {
	return &Logger{writer: w, logType: logType}
}

func NewActual(dir string, t core.LogType) (*Logger, error) {
	var err error
	var file *os.File

	if _, err := os.Stat(dir); err != nil {
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	if file, err = os.OpenFile(path.Join(dir, core.GetActualLogFile()), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666); err != nil {
		return nil, err
	}

	return New(file, t), nil
}

func NewActualFromConfig(cfg core.IConfig) (*Logger, error) {
	return NewActual(cfg.LogPath(), cfg.LogLevel())
}

func (l *Logger) Type() core.LogType {
	return l.logType
}

func (l *Logger) Debugf(format string, args ...any) {
	if (l.logType & core.DebugLog) != 0 {
		fmt.Fprintf(l.writer, "- "+time.Now().Format("15:04:05")+" [DEBUG] "+format+"\n", args...)
	}
}

func (l *Logger) Infof(format string, args ...any) {
	if (l.logType & core.InfoLog) != 0 {
		fmt.Fprintf(l.writer, "- "+time.Now().Format("15:04:05")+" [INFO] "+format+"\n", args...)
	}
}

func (l *Logger) Warnf(format string, args ...any) {
	if (l.logType & core.WarnLog) != 0 {
		fmt.Fprintf(l.writer, "- "+time.Now().Format("15:04:05")+" [WARN] "+format+"\n", args...)
	}
}

func (l *Logger) Errorf(format string, args ...any) {
	if (l.logType & core.ErrLog) != 0 {
		fmt.Fprintf(l.writer, "- "+time.Now().Format("15:04:05")+" [ERROR] "+format+"\n", args...)
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
