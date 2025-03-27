package logger

import (
	"log"
	"os"
	"strings"
)

// Level represents the severity of the log message
type Level int

// log levels
const (
	DEBUG Level = iota // 0
	INFO               // 1
	WARN               // 2
	ERROR              // 3
)

// Logger wraps the standard logger with level-based filtering
type Logger struct {
	*log.Logger
	level Level
}

// NewLogger creates a logger with a prefix and minimum log level
func NewLogger(prefix string, level Level) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, prefix, log.LstdFlags),
		level:  level,
	}
}

// ParseLevel converts a string to a Level
func ParseLevel(levelStr string) (Level, error) {
	switch strings.ToUpper(levelStr) {
	case "DEBUG":
		return DEBUG, nil
	case "WARN":
		return WARN, nil
	case "ERROR":
		return ERROR, nil
	default:
		return INFO, nil
	}
}

// Log logs a message if its level meets the logger's level
func (l *Logger) Log(level Level, format string, v ...any) {
	if level >= l.level {
		l.Printf(format, v...)
	}
}

// Convenience methods for Debug level
func (l *Logger) Debug(format string, v ...any) {
	l.Log(DEBUG, format, v...)
}

// Convenience methods for Info level
func (l *Logger) Info(format string, v ...any) {
	l.Log(INFO, format, v...)
}

// Convenience methods for Warn level
func (l *Logger) Warn(format string, v ...any) {
	l.Log(WARN, format, v...)
}

// Convenience methods for Error level
func (l *Logger) Error(format string, v ...any) {
	l.Log(ERROR, format, v...)
}
