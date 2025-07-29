package writer

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/optiflowic/ghkit/internal/logger"
	"github.com/stretchr/testify/assert"
)

func Test_Write(t *testing.T) {
	tmpDir := t.TempDir()
	log := logger.NewWithWriter(io.Discard, logger.LevelError)
	w := New(log)

	t.Run("success", func(t *testing.T) {
		filePath := filepath.Join(tmpDir, "test", "file.txt")
		content := []byte("hello world")

		err := w.Write(filePath, content)

		assert.NoError(t, err)

		data, readErr := os.ReadFile(filePath)
		assert.NoError(t, readErr)
		assert.Equal(t, content, data)
	})

	t.Run("create directory failure", func(t *testing.T) {
		readonlyDir := filepath.Join(tmpDir, "readonly")
		assert.NoError(t, os.Mkdir(readonlyDir, 0500))

		t.Cleanup(func() {
			_ = os.Chmod(readonlyDir, 0700)
		})

		path := filepath.Join(readonlyDir, "subdir", "file.txt")
		w := New(logger.NewFromFlags(false, false))
		err := w.Write(path, []byte("test"))

		assert.Error(t, err)
	})

	t.Run("write file failure", func(t *testing.T) {
		dirPath := filepath.Join(tmpDir, "asfile")
		assert.NoError(t, os.MkdirAll(dirPath, 0755))

		err := w.Write(dirPath, []byte("data"))

		assert.Error(t, err)
	})
}
