package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jgfranco17/readysetgo/core/pkg/build"
	"github.com/jgfranco17/readysetgo/core/pkg/models"
	"github.com/sirupsen/logrus"
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

func GenerateProject(log *logrus.Entry, projectName string) error {
	// Create project directory
	err := build.BuildGoAppRepository(log, projectName)
	if err != nil {
		return fmt.Errorf("Failed to create base Go project '%s': %v", projectName, err)
	}

	// Create basic project files
	for _, baseFile := range ROOT_LEVEL_FILES {
		_, err = os.Create(baseFile)
		if err != nil {
			return fmt.Errorf("Failed to create file '%s': %v", baseFile, err)
		}
		log.Debugf("Created '%s' file..", baseFile)
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

	fmt.Printf("Created '%s' project!", projectName)
	return nil
}
