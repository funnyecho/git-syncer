package runners

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

// PrintUsageErr try to print usage info from flagset, and if match, error will be swallowed
func PrintUsageErr(_ []string) (runner.BubbleTap, error) {
	return func(e error) error {
		if flagset, flagsetErr := UseFlagset(); flagsetErr == nil {
			errCode := errors.GetErrorCode(e)
			if errCode == exitcode.Usage || errCode == exitcode.MissingArguments {
				flagset.Usage()
				return nil
			}
		}

		return e
	}, nil
}
