package gitter

import (
	"github.com/funnyecho/git-syncer/pkg/command"
)

func GetProjectDir() (dir string, err error) {
	cmd := command.Command("git", "rev-parse --show-toplevel")

	r, err := cmd.Output()
	if err != nil {
		return
	}

	dir = string(r)
	return
}
