package gitter

func (g *git) GetHead() (string, error) {
	cmd := g.command("git", "rev-parse", "HEAD")

	v, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(v), nil
}
