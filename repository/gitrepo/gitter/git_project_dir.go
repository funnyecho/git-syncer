package gitter

func (g *git) GetProjectDir() (dir string, err error) {
	cmd := g.command("git", "rev-parse", "--show-toplevel")

	r, err := cmd.Output()
	if err != nil {
		return
	}

	dir = string(r)
	return
}
