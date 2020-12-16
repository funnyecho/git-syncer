package gitter

import "github.com/funnyecho/git-syncer/pkg/command"

func Checkout(head string) error {
	cmd := command.Command("git", "checkout", head)

	return cmd.Run()
}
