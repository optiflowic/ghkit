package logger

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Info(t *testing.T) {
	buf := new(bytes.Buffer)
	level := slog.LevelInfo
	log := NewWithWriter(buf, level)
	msg := "Info message"

	log.Info(msg)

	out := buf.String()
	assert.Contains(t, out, "[INFO]")
	assert.Contains(t, out, msg)
}

func Test_Error(t *testing.T) {
	buf := new(bytes.Buffer)
	level := slog.LevelInfo
	log := NewWithWriter(buf, level)
	msg := "Error message"

	log.Error(msg)

	out := buf.String()
	assert.Contains(t, out, "[ERROR]")
	assert.Contains(t, out, msg)
}

func Test_DebugDisabled(t *testing.T) {
	buf := new(bytes.Buffer)
	level := slog.LevelInfo
	log := NewWithWriter(buf, level)
	msg := "Debug message"

	log.Debug(msg)

	assert.Empty(t, buf)
}

func Test_DebugEnabled(t *testing.T) {
	buf := new(bytes.Buffer)
	level := slog.LevelDebug
	log := NewWithWriter(buf, level)
	msg := "Debug message"

	log.Debug(msg)

	out := buf.String()
	assert.Contains(t, out, "[DEBUG]")
	assert.Contains(t, out, msg)
}
