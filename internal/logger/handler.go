package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

type Handler struct {
	out   io.Writer
	level slog.Level
}

func NewHandler(out io.Writer, level slog.Level) *Handler {
	return &Handler{
		out:   out,
		level: level,
	}
}

func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	ts := r.Time.Format("2006-01-02 15:04:05")
	attrs := []string{}
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a.String())
		return true
	})
	_, err := fmt.Fprintf(h.out, "%s [%s] %s (%s)", ts, r.Level.String(), r.Message, strings.Join(attrs, ", "))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(h.out)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) WithAttrs(_ []slog.Attr) slog.Handler { return h }
func (h *Handler) WithGroup(_ string) slog.Handler      { return h }
