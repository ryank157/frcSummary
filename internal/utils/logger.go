package utils

import (
	"log"
	"os"
)

// Logger struct.  For simplicity, this uses the standard `log` package.  For a more complex application,
// consider using a more advanced logger like `zap` or `logrus`.
type Logger struct {
	level string
	log   *log.Logger
}

// NewLogger creates a new logger with the given log level.
func NewLogger(level string) *Logger {
	return &Logger{
		level: level,
		log:   log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Infof logs an info message.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level == "debug" || l.level == "info" {
		l.log.Printf("[INFO] "+format, v...)
	}
}

// Debugf logs a debug message.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level == "debug" {
		l.log.Printf("[DEBUG] "+format, v...)
	}

}

// Warnf logs a warning message.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.log.Printf("[WARN] "+format, v...)

}

// Errorf logs an error message.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log.Printf("[ERROR] "+format, v...)
}

// Fatalf logs a fatal message and exits.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log.Fatalf("[FATAL] "+format, v...)
}
