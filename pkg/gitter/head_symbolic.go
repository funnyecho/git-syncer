package gitter

import (
	"github.com/funnyecho/git-syncer/pkg/command"
	"strings"
)

func GetSymbolicHead() string {
	cmd := command.Command("git", "symbolic-ref", "HEAD")

	v, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimPrefix(string(v), "refs/heads/")
}
