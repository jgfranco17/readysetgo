package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jgfranco17/readysetgo/core/pkg/models"
)

var ROOT_LEVEL_FILES = []string{"main.go", "README.md"}
var project = models.Project{}

func fillDirectory(directory string, files []string) error {
	for _, file := range files {
		filePath := filepath.Join(directory, file)
		_, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("Failed to fill directory '%s': %v", directory, err.Error())
		}
	}
	return nil
}

func GenerateProject(projectName string) error {
	// Create project directory
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		return fmt.Errorf("Failed to create project '%s': %v", projectName, err.Error())
	}

	// Create basic project files
	for _, file := range ROOT_LEVEL_FILES {
		filePath := filepath.Join(projectName, file)
		_, err := os.Create(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
