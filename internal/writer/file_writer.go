package writer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/optiflowic/ghkit/internal/logger"
)

type FileWriter struct {
	log logger.Logger
}

func New(log logger.Logger) FileWriter {
	return FileWriter{log: log}
}

func (w FileWriter) Write(path string, data []byte) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0700); err != nil {
		w.log.Error("Failed to create directories", "path", dir, "error", err)
		return fmt.Errorf("failed to create directories: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		w.log.Error("Failed to write file", "path", path, "error", err)
		return fmt.Errorf("failed to write file: %w", err)
	}

	w.log.Info("File written", "path", path, "bytes", len(data))
	return nil
}
