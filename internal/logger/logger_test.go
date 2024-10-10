package logger

import (
	"bytes"
	"log"
	"testing"
)

func TestLoggerOutputs(t *testing.T) {
	var buf bytes.Buffer
	testLogger := &KaizeenLogger{logger: log.New(&buf, "", log.LstdFlags)}

	testLogger.Info("This is an info message.")
	if !bytes.Contains(buf.Bytes(), []byte("INFO: This is an info message.")) {
		t.Errorf("Expected Info log entry to be written.")
	}

	buf.Reset()
	testLogger.Error("This is an error message.")
	if !bytes.Contains(buf.Bytes(), []byte("ERROR: This is an error message.")) {
		t.Errorf("Expected Error log entry to be written.")
	}
}
