package build

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Init a list of Go modules
func initModule(mainDir string, names []string) error {
	var failedDirectories []string
	for _, name := range names {
		// Make directory and add 'pkg' subdir inside
		subdir := filepath.Join(mainDir, name)
		err := os.Mkdir(subdir, 0755)
		if err != nil {
			failedDirectories = append(failedDirectories, subdir)
		}
		subdirPkg := filepath.Join(subdir, "pkg")
		err = os.Mkdir(subdirPkg, 0755)
		if err != nil {
			failedDirectories = append(failedDirectories, subdirPkg)
		}
		cmd := exec.Command("go", "mod", "init", fmt.Sprintf("./%s", name))
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("Error initializing Go module: %v", err)
		}
	}
	return nil
}

// Create a new Go app repository
func BuildGoAppRepository(log *logrus.Entry, projectName string) error {
	// Specify the name of the repository and workspace
	workspaceName := "src"
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Failed to identify current working directory: %v", err)
	}
	projectPath := filepath.Join(cwd, projectName)

	// Create the root directory for the repository
	err = os.Mkdir(projectPath, 0755)
	if err != nil {
		return fmt.Errorf("Error creating repository directory: %v", err)
	}

	// Create the workspace directory inside the repository
	workspacePath := filepath.Join(projectPath, workspaceName)
	err = os.Mkdir(workspacePath, 0755)
	if err != nil {
		return fmt.Errorf("Error creating workspace directory: %v", err)
	}

	// Run "go mod init" command inside the workspace
	err = initModule(projectPath, []string{})
	if err != nil {
		return fmt.Errorf("Error initializing workspace modules: %v", err)
	}

	// Create a sample Go file inside the workspace
	sampleFilePath := filepath.Join(workspacePath, "main.go")
	sampleFileContent := `package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
`
	err = os.WriteFile(sampleFilePath, []byte(sampleFileContent), 0644)
	if err != nil {
		return fmt.Errorf("Error creating sample Go file: %v", err)
	}

	fmt.Println("Go repository with workspace created successfully.")
	return nil
}
