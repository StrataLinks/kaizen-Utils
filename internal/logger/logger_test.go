package logger

import (
	"testing"
)

// TestLogHandler is a mock handler for testing.
type TestLogHandler struct {
	loggedMessages []string
}

func (h *TestLogHandler) Log(level LogLevel, message string) {
	h.loggedMessages = append(h.loggedMessages, message)
}

func TestLoggerOutputs(t *testing.T) {
	handler := &TestLogHandler{}
	logger := NewLogger([]LogHandle{handler})

	logger.Info("This is an info message.")
	expectedMessage := "This is an info message."
	if handler.loggedMessages[0] != expectedMessage {
		t.Errorf("Expected '%s', got '%s'", expectedMessage, handler.loggedMessages[0])
	}
}
