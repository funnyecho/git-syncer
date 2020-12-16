package gitter

import (
	"github.com/funnyecho/git-syncer/pkg/command"
)

func GetHead() string {
	cmd := command.Command("git", "rev-parse", "HEAD")

	v, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(v)
}