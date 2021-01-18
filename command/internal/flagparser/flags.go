package flagparser

import (
	"flag"

	"github.com/funnyecho/git-syncer/pkg/flagbinder"
)

// Parser wrap flags parsing
func Parser(options interface{}, flags *flag.FlagSet) func(args []string) error {
	bindFlagErr := flagbinder.Bind(&options, flags)
	if bindFlagErr != nil {
		return func(args []string) error {
			return bindFlagErr
		}
	}

	return func(args []string) error {
		return flags.Parse(args)
	}
}
