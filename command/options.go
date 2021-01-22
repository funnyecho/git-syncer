package command

import (
	"flag"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/flagbinder"
	"github.com/funnyecho/git-syncer/pkg/log"
)

// BasicOptions basic options for command
type BasicOptions struct {
	Base   string `flag:"base" value:"" usage:"Base dir path to run syncer"`
	Branch string `flag:"branch" value:"" usage:"Push a specific branch"`
	Remote string `flag:"remote" value:"" usage:"Push to specific remote"`
}

type OptionsParser func(args []string) int

func OptionsBinder(options interface{}, flags *flag.FlagSet) OptionsParser {
	bindFlagErr := flagbinder.Bind(options, flags)
	if bindFlagErr != nil {
		return func(args []string) int {
			log.Errore("failed to bind arguments", bindFlagErr)
			return exitcode.Usage
		}
	}

	return func(args []string) int {
		if err := flags.Parse(args); err != nil {
			log.Errore("failed to parse arguments", bindFlagErr)
			return exitcode.Usage
		}

		return exitcode.Nil
	}
}
