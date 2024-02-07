package gitclient

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func InitializeGitRepo(path string) error {
	cmd := exec.Command("git", "init", ".")
	err := cmd.Run()
	if err != nil {
		errorMessge := fmt.Sprintf("Error initializing Git repository: %v", err.Error())
		log.Error(errorMessge)
		return fmt.Errorf(errorMessge)
	}
	return nil
}
