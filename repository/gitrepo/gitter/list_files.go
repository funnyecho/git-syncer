package gitter

import (
	"strings"
)

func (g *git) ListFiles(path string) ([]string, error) {
	if path == "" {
		path = "."
	}
	cmd := g.command("git", "ls-files", "-z", "--", path)

	v, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(string(v), ""), nil
}
