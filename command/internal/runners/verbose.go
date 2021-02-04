package runners

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/variable"
)

// Verbose set verbose to variable
func Verbose(_ []string) (runner.BubbleTap, error) {
	if v, vErr := useFlag(flagVerbose); vErr == nil {
		variable.WithVerbose(v)
	}

	return nil, nil
}
