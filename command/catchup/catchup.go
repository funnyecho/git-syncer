package catchup

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/command/internal/runners"
	"github.com/funnyecho/git-syncer/syncer"
	"github.com/mitchellh/cli"
)

// Factory command `catchup` factory
func Factory() (cli.Command, error) {
	return &cmd{}, nil
}

type cmd struct {
}

func (c *cmd) Help() string {
	return "Creates or updates the log file on the remote.\n" +
		"It assumes that you uploaded all other files already.\n" +
		"You might have done that with another program."
}

func (c *cmd) Synopsis() string {
	return "Creates or updates the log file on the remote."
}

func (c *cmd) Run(args []string) (ext int) {
	opt := &options{}

	return runner.Run(
		args,
		runners.PrintErr,
		runners.PrintUsageErr,
		runners.WithFlagset("catchup", opt),
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

			return syncer.Catchup(remote, gitter)
		}),
	)
}
