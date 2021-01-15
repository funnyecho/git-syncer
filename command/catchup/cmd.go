package catchup

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/repository"
	"github.com/mitchellh/cli"
)

// Factory of command `catchup`
func Factory() (cli.Command, error) {
	return &cmd{}, nil
}

type cmd struct {
}

func (c *cmd) Help() string {
	return "Creates or updates the `.git-ftp.log` file on the remote contrib.\n" +
		"It assumes that you uploaded all other files already.\n" +
		"You might have done that with another program."
}

func (c *cmd) Synopsis() string {
	return "Creates or updates the `.git-ftp.log` file on the remote contrib."
}

func (c *cmd) Run(args []string) int {
	return runner.Run("catchup", args, runner.WithTapCommand(func(r repository.Repository, c contrib.Contrib) error {
		return Catchup(c, r)
	}))
}
