package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jgfranco17/readysetgo/core/pkg/models"
	log "github.com/sirupsen/logrus"
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

func fillDirectory(content models.Subdirectory) error {
	for _, file := range content.Files {
		filePath := filepath.Join(content.Name, file.FullName())
		_, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("Failed to fill directory '%s' with %d files: %v", content.Name, len(content.Files), err)
		}
	}
	return nil
}

func GenerateProject(projectName string) error {
	// Create project directory
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		return fmt.Errorf("Failed to create project '%s': %v", projectName, err)
	}

	// Create basic project files
	for _, baseFile := range ROOT_LEVEL_FILES {
		_, err := os.Create(baseFile)
		if err != nil {
			return fmt.Errorf("Failed to create file '%s': %v", baseFile, err)
		}
	}

	// Create project subdirectories
	for _, subdir := range projectSubdirs {
		dirPath := filepath.Join(projectName, subdir.Name)
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("Failed to create subdirectory '%s': %v", subdir.Name, err)
		}
		fillDirectory(subdir)
	}

	log.Debugf("")
	return nil
}
