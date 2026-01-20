package util

import (
	"log"
	"os"
	"time"
)

// Logger provides structured logging
type Logger struct {
	*log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

// LogError logs error with context
func (l *Logger) LogError(err error, context map[string]interface{}) {
	l.Printf("ERROR: %v | Context: %+v", err, context)
}

// LogInfo logs info message
func (l *Logger) LogInfo(message string, context map[string]interface{}) {
	l.Printf("INFO: %s | Context: %+v", message, context)
}

// LogRequest logs incoming request details
func (l *Logger) LogRequest(method, path, clientIP string, duration time.Duration) {
	l.Printf("REQUEST: %s %s | IP: %s | Duration: %v", method, path, clientIP, duration)
}
