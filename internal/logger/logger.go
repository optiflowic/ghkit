package logger

import (
	"io"
	"log/slog"
	"os"
)

type Logger struct {
	slog *slog.Logger
}

func New(level *Level) Logger {
	l := LevelInfo
	if level != nil {
		l = *level
	}
	return NewWithWriter(os.Stdout, l)
}

func NewFromFlags(verbose, debug bool) Logger {
	var level Level
	switch {
	case debug:
		level = LevelDebug
	case verbose:
		level = LevelInfo
	default:
		level = LevelError
	}

	return NewWithWriter(os.Stdout, level)
}

func NewWithWriter(w io.Writer, level Level) Logger {
	handler := NewHandler(w, level.toSlogLevel())
	return Logger{slog: slog.New(handler)}
}

func (l *Logger) Info(msg string, args ...any)  { l.slog.Info(msg, args...) }
func (l *Logger) Debug(msg string, args ...any) { l.slog.Debug(msg, args...) }
func (l *Logger) Warn(msg string, args ...any)  { l.slog.Warn(msg, args...) }
func (l *Logger) Error(msg string, args ...any) { l.slog.Error(msg, args...) }
