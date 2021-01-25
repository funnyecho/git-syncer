package gitter

import "strings"

func (g *git) GetHeadSHA1() (string, error) {
	cmd := g.command("git", "log", "-n 1", "--pretty=format:%H")

	v, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(v)), nil
}
