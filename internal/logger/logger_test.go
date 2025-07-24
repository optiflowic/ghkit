package logger

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	msg := "message"

	log := New(nil)
	log.Info(msg)

	w.Close()
	os.Stdout = oldStdout
	var output bytes.Buffer
	_, _ = output.ReadFrom(r)

	assert.Contains(t, output.String(), "[INFO]")
	assert.Contains(t, output.String(), msg)
}


func Test_Info(t *testing.T) {
	buf := new(bytes.Buffer)
	level := LevelInfo
	log := NewWithWriter(buf, level)
	msg := "Info message"

	log.Info(msg)

	out := buf.String()
	assert.Contains(t, out, "[INFO]")
	assert.Contains(t, out, msg)
}

func Test_Error(t *testing.T) {
	buf := new(bytes.Buffer)
	level := LevelInfo
	log := NewWithWriter(buf, level)
	msg := "Error message"

	log.Error(msg)

	out := buf.String()
	assert.Contains(t, out, "[ERROR]")
	assert.Contains(t, out, msg)
}

func Test_DebugDisabled(t *testing.T) {
	buf := new(bytes.Buffer)
	level := LevelInfo
	log := NewWithWriter(buf, level)
	msg := "Debug message"

	log.Debug(msg)

	assert.Empty(t, buf)
}

func Test_DebugEnabled(t *testing.T) {
	buf := new(bytes.Buffer)
	level := LevelDebug
	log := NewWithWriter(buf, level)
	msg := "Debug message"

	log.Debug(msg)

	out := buf.String()
	assert.Contains(t, out, "[DEBUG]")
	assert.Contains(t, out, msg)
}
