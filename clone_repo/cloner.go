package clone_repo

import (
	"os"
	"os/exec"
)

func CloneRepo(url string, dir string) error {
	// Clone the repository using Git
	cmd := exec.Command("git", "clone", url, dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
