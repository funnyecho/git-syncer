package gitter

func (g *git) Checkout(head string) error {
	cmd := g.command("git", "checkout", head)

	return cmd.Run()
}
