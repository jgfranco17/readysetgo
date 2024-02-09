package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jgfranco17/readysetgo/core/pkg/models"
)

var ROOT_LEVEL_FILES = []string{"main.go", "README.md"}
var projectSubdirs = []models.Subdirectory{
	{
		Name: ".github",
		Files: []models.File{
			{Name: "actions", Extension: "yaml"},
		},
	},
}

func fillDirectory(content *models.Subdirectory) error {
	for _, file := range content.Files {
		filePath := filepath.Join(content.Name, file.FullName())
		_, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("Failed to fill directory '%s' with %d files: %v", content.Name, len(content.Files), err)
		}
	}
	return nil
}

func GenerateProject(projectName string, minimal bool) error {
	// Create project directory
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		return fmt.Errorf("Failed to create project '%s': %v", projectName, err)
	}

	// Create basic project files
	for _, file := range ROOT_LEVEL_FILES {
		_, err := os.Create(file)
		if err != nil {
			return fmt.Errorf("Failed to create file '%s': %v", file, err)
		}
	}

	return nil
}
