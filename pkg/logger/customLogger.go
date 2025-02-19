package logger

import (
	"log"
	"os"
)

const (
	INFO  = "INFO"
	DEBUG = "DEBUG"
)

var instance *Logger

type Logger struct {
	infoEnabled  bool
	debugEnabled bool
	errorEnabled bool
	errorLogger  *log.Logger
	infoLogger   *log.Logger
	debugLogger  *log.Logger
}

func NewLogger(level string) *Logger {
	logger := &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime),
	}

	switch level {
	case INFO:
		logger.infoEnabled = true
		logger.errorEnabled = true
	case DEBUG:
		logger.infoEnabled = true
		logger.debugEnabled = true
		logger.errorEnabled = true
	}

	// Set the package-level logger instance
	instance = logger

	return logger
}

// Info Static access to the logger instance
func Info(msg string) {
	if instance != nil && instance.infoEnabled {
		instance.infoLogger.Println(msg)
	}
}

func Error(msg string) {
	if instance != nil && instance.errorEnabled {
		instance.infoLogger.Println(msg)
	}
}

func Debug(msg string) {
	if instance != nil && instance.debugEnabled {
		instance.debugLogger.Println(msg)
	}
}
