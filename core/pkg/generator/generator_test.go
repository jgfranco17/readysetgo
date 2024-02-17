package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/jgfranco17/readysetgo/core/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestFillDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Define some sample file names
	files := []models.File{
		{Name: "file1", Extension: "txt"},
		{Name: "file2", Extension: "txt"},
	}

	subdir := models.Subdirectory{Name: tempDir, Files: files}
	err := fillDirectory(subdir)
	assert.NoError(t, err, "Failed to fill subdirectories")

	for _, file := range files {
		filePath := filepath.Join(tempDir, file.FullName())
		_, err := os.Stat(filePath)
		assert.NoError(t, err, fmt.Sprintf("Expected file '%s' to exist", filePath))

		// Clean up created files
		err = os.Remove(filePath)
		assert.NoError(t, err, fmt.Sprintf("Failed to remove file '%s'", filePath))
	}

	err = os.RemoveAll(tempDir)
	assert.NoError(t, err, "Failed to delete temp dir")
}
