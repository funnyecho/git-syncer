package git

import (
	"os/exec"
	"strings"
)

// New new git
func New() *Git {
	return &Git{}
}

// Git gitter implementation
type Git struct {
}

func output(args []string) (string, error) {
	cmd := exec.Command("git", args...)

	v, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(v)), nil
}
