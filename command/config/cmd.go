package config

import (
	"flag"
	"fmt"

	"github.com/funnyecho/git-syncer/command/internal/flagparser"
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/repository"
	"github.com/mitchellh/cli"
)

// Factory of command `push`
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

func (c *cmd) Run(args []string) int {
	return runner.Run("config", args, runner.WithTapCommand(func(r repository.Repository, c contrib.Contrib) error {
		var options Options

		flags := flag.NewFlagSet("config", flag.ContinueOnError)

		if argsErr := flagparser.Parser(&options, flags)(args); argsErr != nil {
			log.Errore("Failed to parse config flags", argsErr)
			return errors.NewError(
				errors.WithCode(exitcode.Usage),
				errors.WithErr(argsErr),
			)
		}

		v := flags.Args()
		if len(v) == 0 {
			return errors.NewError(
				errors.WithMsg("failed to get key or value"),
				errors.WithCode(exitcode.Usage),
			)
		}

		if len(v) == 1 {
			value, configErr := GetConfig(v[0], r)
			if configErr != nil {
				return errors.NewError(
					errors.WithMsgf("failed to get config: %s", v[0]),
					errors.WithErr(configErr),
				)
			}
			fmt.Println(value)
			return nil
		}

		return UpdateConfig(v[0], v[1], r)
	}))
}
