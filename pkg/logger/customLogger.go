package logger

import (
	"log"
	"os"
)

const (
	INFO  = "INFO"
	DEBUG = "DEBUG"
)

type Logger struct {
	infoEnabled  bool
	debugEnabled bool
	infoLogger   *log.Logger
	debugLogger  *log.Logger
}

func NewLogger(level string) *Logger {

	logger := &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime),
	}

	switch level {
	case INFO:
		logger.infoEnabled = true
	case DEBUG:
		logger.infoEnabled = true
		logger.debugEnabled = true
	}

	return logger
}

func (l *Logger) Info(msg string) {
	if l.infoEnabled {
		l.infoLogger.Println(msg)
	}
}

func (l *Logger) Debug(msg string) {
	if l.debugEnabled {
		l.debugLogger.Println(msg)
	}
}
