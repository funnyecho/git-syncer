package setup

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/command/internal/runners"
	"github.com/funnyecho/git-syncer/syncer"
	"github.com/mitchellh/cli"
)

// Factory command `setup` factory
func Factory() (cli.Command, error) {
	return &cmd{}, nil
}

type cmd struct {
}

func (c *cmd) Help() string {
	return "Uploads all git-tracked non-ignored files to the remote contrib and " +
		"creates the log file containing the SHA1 of the latest commit."
}

func (c *cmd) Synopsis() string {
	return "Setup remote contrib to the latest commit of repo"
}

func (c *cmd) Run(args []string) (ext int) {
	opt := &options{}

	return runner.Run(
		args,
		runners.PrintErr,
		runners.PrintUsageErr,
		runners.WithFlagset("setup", opt),
		runners.WorkingDir,
		runners.Gitter,
		runners.Remote,
		runners.WithTarget(func(_ []string) error {
			gitter, gitterErr := runners.UseGitter()
			if gitterErr != nil {
				return gitterErr
			}

			remote, remoteErr := runners.UseRemote()
			if remoteErr != nil {
				return remoteErr
			}

			return syncer.Setup(remote, gitter)
		}),
	)
}
