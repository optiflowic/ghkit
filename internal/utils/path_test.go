package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Exists(t *testing.T) {
	t.Run("file exists", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "testfile-*.txt")
		assert.NoError(t, err)
		defer func() {
			err := os.Remove(tmpFile.Name())
			assert.Nil(t, err)
		}()

		assert.Equal(t, true, Exists(tmpFile.Name()))
	})

	t.Run("directory exists", func(t *testing.T) {
		tmpDir := t.TempDir()
		assert.Equal(t, true, Exists(tmpDir))
	})

	t.Run("file does not exist", func(t *testing.T) {
		nonExistentPath := filepath.Join(os.TempDir(), "nonexistent-file.txt")
		assert.Equal(t, false, Exists(nonExistentPath))
	})
}
