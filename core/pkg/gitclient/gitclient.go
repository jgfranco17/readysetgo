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
		return fmt.Errorf("Error initializing Git repository: %v", err)
	}
	log.Info("Git repository created.")
	return nil
}
