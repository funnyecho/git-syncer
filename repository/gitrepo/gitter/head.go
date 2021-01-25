package gitter

import "strings"

func (g *git) GetHead() (string, error) {
	cmd := g.command("git", "rev-parse", "HEAD")

	v, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(v)), nil
}
