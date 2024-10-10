package logger

import (
	"sync"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

type LogHandle interface {
	Log(level LogLevel, message string)
}

type Logger struct {
	handle []LogHandle
	mu     sync.Mutex
}

func NewLogger(handle []LogHandle) *Logger {
	return &Logger{
		handle: handle,
	}
}

func (l *Logger) log(level LogLevel, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, handler := range l.handle {
		handler.Log(level, message)
	}
}

// Info logs a message at Info level.
func (l *Logger) Info(message string) {
	l.log(InfoLevel, message)
}

// Debug logs a message at Debug level.
func (l *Logger) Debug(message string) {
	l.log(DebugLevel, message)
}

// Warn logs a message at Warn level.
func (l *Logger) Warn(message string) {
	l.log(WarnLevel, message)
}

// Error logs a message at Error level.
func (l *Logger) Error(message string) {
	l.log(ErrorLevel, message)
}

// AddHandler adds a new log handle to the logger.
func (l *Logger) AddHandler(handle LogHandle) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.handle = append(l.handle, handle)
}
