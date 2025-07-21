package logger

import (
	"io"
	"log/slog"
	"os"
)

type Logger struct {
	slog *slog.Logger
}

func New(level *slog.Level) *Logger {
	l := slog.LevelInfo
	if level != nil {
		l = *level
	}
	return NewWithWriter(os.Stdout, l)
}

func NewWithWriter(w io.Writer, level slog.Level) *Logger {
	handler := NewHandler(w, level)
	return &Logger{slog: slog.New(handler)}
}

func (l *Logger) Info(msg string, args ...any)  { l.slog.Info(msg, args...) }
func (l *Logger) Debug(msg string, args ...any) { l.slog.Debug(msg, args...) }
func (l *Logger) Error(msg string, args ...any) { l.slog.Error(msg, args...) }
