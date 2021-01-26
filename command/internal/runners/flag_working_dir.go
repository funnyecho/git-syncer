package runners

import (
	"os"

	"github.com/funnyecho/git-syncer/command/internal/runner"
)

// WorkingDir sets working dir
func WorkingDir(args []string) (runner.BubbleTap, error) {
	v, err := useFlag(flagWorkingDir)
	if v == "" {
		return nil, err
	}

	err = os.Chdir(v)
	return nil, err
}
