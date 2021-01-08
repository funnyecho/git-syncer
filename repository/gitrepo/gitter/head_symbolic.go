package gitter

import (
	"strings"
)

func (g *git) GetSymbolicHead() string {
	cmd := g.command("git", "symbolic-ref", "HEAD")

	v, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimPrefix(string(v), "refs/heads/")
}