package build

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestBuildGoAppRepository(t *testing.T) {
	log := &logrus.Entry{}
	projectName := "test_build_repo"

	// Clean up test directory after the test
	defer func() {
		err := os.RemoveAll(projectName)
		if err != nil {
			fmt.Printf("Error cleaning up test directory: %v\n", err)
		}
	}()

	err := BuildGoAppRepository(log, projectName)
	assert.NoError(t, err, "Unexpected error in BuildGoAppRepository")

	// Verify that repository directory and sample Go file were created
	expectedItems := []string{projectName, fmt.Sprintf("%s/src/main.go", projectName)}
	for _, item := range expectedItems {
		_, err = os.Stat(item)
		assert.NoError(t, err, "Expected '%s' to exist", item)
	}
}
