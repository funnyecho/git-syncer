package gitter

import (
	"strings"
)

func (g *git) GetSymbolicHead() (string, error) {
	cmd := g.command("git", "symbolic-ref", "HEAD")

	v, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(strings.TrimSpace(string(v)), "refs/heads/"), nil
}
