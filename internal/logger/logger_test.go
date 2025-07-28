package logger

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	level := LevelError
	tests := []struct {
		name      string
		level     *Level
		msgPrefix string
		msg       string
	}{
		{name: "Do not set level", level: nil, msgPrefix: "[INFO]", msg: "message"},
		{name: "Setting the level", level: &level, msgPrefix: "", msg: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			msg := "message"

			log := New(tt.level)
			log.Info(msg)

			err := w.Close()
			os.Stdout = oldStdout
			var output bytes.Buffer
			_, _ = output.ReadFrom(r)
			outputString := output.String()

			assert.Nil(t, err)
			assert.Contains(t, outputString, tt.msgPrefix)
			assert.Contains(t, outputString, tt.msg)
		})
	}
}

func Test_NewFromFlags(t *testing.T) {
	tests := []struct {
		name             string
		verbose          bool
		debug            bool
		includedPrefixes []string
		excludedPrefixes []string
	}{
		{
			name:             "Do not set flags",
			verbose:          false,
			debug:            false,
			includedPrefixes: []string{"[ERROR]"},
			excludedPrefixes: []string{"[DEBUG]", "[INFO]", "[WARN]"},
		},
		{
			name:             "Set the verbose flag",
			verbose:          true,
			debug:            false,
			includedPrefixes: []string{"[INFO]", "[WARN]", "[ERROR]"},
			excludedPrefixes: []string{"[DEBUG]"},
		},
		{
			name:             "Set the debug flag",
			verbose:          false,
			debug:            true,
			includedPrefixes: []string{"[DEBUG]", "[INFO]", "[WARN]", "[ERROR]"},
			excludedPrefixes: []string{},
		},
		{
			name:             "Set the verbose and debug flag",
			verbose:          true,
			debug:            true,
			includedPrefixes: []string{"[DEBUG]", "[INFO]", "[WARN]", "[ERROR]"},
			excludedPrefixes: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			msg := "message"

			log := NewFromFlags(tt.verbose, tt.debug)
			log.Debug(msg)
			log.Info(msg)
			log.Warn(msg)
			log.Error(msg)

			err := w.Close()
			os.Stdout = oldStdout
			var output bytes.Buffer
			_, _ = output.ReadFrom(r)
			outputString := output.String()

			assert.Nil(t, err)
			for _, prefix := range tt.includedPrefixes {
				assert.Contains(t, outputString, prefix)
			}
			for _, prefix := range tt.excludedPrefixes {
				assert.NotContains(t, outputString, prefix)
			}
		})
	}
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

func Test_Warn(t *testing.T) {
	buf := new(bytes.Buffer)
	level := LevelInfo
	log := NewWithWriter(buf, level)
	msg := "Warn message"

	log.Warn(msg)

	out := buf.String()
	assert.Contains(t, out, "[WARN]")
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
