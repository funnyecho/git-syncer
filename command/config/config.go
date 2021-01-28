package config

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/command/internal/runners"
	"github.com/funnyecho/git-syncer/syncer"
	"github.com/mitchellh/cli"
)

// Factory of command config
func Factory() (cli.Command, error) {
	return &cmd{}, nil
}

type cmd struct {
}

func (c *cmd) Help() string {
	return "Config getter and setter"
}

func (c *cmd) Synopsis() string {
	return "Config getter and setter"
}

func (c *cmd) Run(args []string) (ext int) {
	opt := &Options{}

	return runner.Run(
		args,
		runners.PrintErr,
		runners.PrintUsageErr,
		runners.WithFlagset("config", opt),
		runners.WorkingDir,
		runners.Gitter,
		runners.WorkingHead,
		runners.WithTarget(func(_ []string) error {
			flagset, flagsetErr := runners.UseFlagset()
			if flagsetErr != nil {
				return flagsetErr
			}

			gitter, gitterErr := runners.UseGitter()
			if gitterErr != nil {
				return gitterErr
			}

			return syncer.Config(gitter, flagset.Args())
		}),
	)
}
