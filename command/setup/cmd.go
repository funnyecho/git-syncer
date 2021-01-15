package setup

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/repository"

	"github.com/mitchellh/cli"
)

// Factory of command `setup`
func Factory() (cli.Command, error) {
	return &cmd{}, nil
}

type cmd struct {
}

func (c *cmd) Help() string {
	return "Uploads all git-tracked non-ignored files to the remote contrib and " +
		"creates the `.git-syncer.log` file containing the SHA1 of the latest commit."
}

func (c *cmd) Synopsis() string {
	return "Setup remote contrib to the latest commit of repo"
}

func (c *cmd) Run(args []string) (ext int) {
	return runner.Run("setup", args, runner.WithTapCommand(func(r repository.Repository, c contrib.Contrib) error {
		return Setup(c, r)
	}))
}
