package gitter

import (
	"github.com/funnyecho/git-syncer/pkg/command"
	"strings"
)

func ListFiles(path string) ([]string, error) {
	if path == "" {
		path = "."
	}
	cmd := command.Command("git", "ls-files", "-z", "--", path)

	v, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(string(v), ""), nil
}
