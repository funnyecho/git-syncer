package runners

import (
	"flag"

	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/flagbinder"
	"github.com/funnyecho/git-syncer/pkg/log"
)

const (
	flagWorkingDir = "wd"
	flagRemote     = "remote"
	flagHead       = "head"
)

// WithFlagset wrap flagset runner
func WithFlagset(name string, opt interface{}) runner.CaptureTap {
	flags = flag.NewFlagSet(name, flag.ContinueOnError)
	fp := flagsBinder(opt, flags)

	return func(args []string) (runner.BubbleTap, error) {
		parseErr := fp(args)
		if parseErr != nil {
			return nil, parseErr
		}
		return nil, nil
	}
}

// UseFlagset get flagset
func UseFlagset() (*flag.FlagSet, error) {
	if flags == nil {
		log.Infow("try to use uninit flagset")
		return nil, errors.Err(exitcode.InvalidRunnerDependency, "failed to get flagset from runner")
	}

	return flags, nil
}

var flags *flag.FlagSet

type flagsParser func(args []string) error

func flagsBinder(options interface{}, flags *flag.FlagSet) flagsParser {
	bindFlagErr := flagbinder.Bind(options, flags)
	if bindFlagErr != nil {
		return func(args []string) error {
			return errors.WrapC(bindFlagErr, exitcode.Usage, "failed to bind flags")
		}
	}

	return func(args []string) error {
		if err := flags.Parse(args); err != nil {
			return errors.WrapC(err, exitcode.Usage, "failed to parse arguments")
		}

		return nil
	}
}

func useFlag(name string) (string, error) {
	flagset, flagsetErr := UseFlagset()
	if flagsetErr != nil {
		return "", flagsetErr
	}

	flag := flagset.Lookup(name)
	if flag != nil {
		if val := flag.Value.String(); val != "" {
			return val, nil
		}
	}

	return "", nil
}
