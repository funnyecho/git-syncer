package options

import (
	"flag"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/flagbinder"
	"github.com/funnyecho/git-syncer/pkg/log"
)

// Parser flag argument parser type
type Parser func(args []string) int

// Binder bind options type with flagset, return argument parser
func Binder(options interface{}, flags *flag.FlagSet) Parser {
	bindFlagErr := flagbinder.Bind(options, flags)
	if bindFlagErr != nil {
		return func(args []string) int {
			log.Errore("failed to bind command flags", bindFlagErr)
			return exitcode.Usage
		}
	}

	return func(args []string) int {
		if err := flags.Parse(args); err != nil {
			log.Errore("failed to parse command flags", bindFlagErr)
			return exitcode.Usage
		}

		return exitcode.Nil
	}
}
