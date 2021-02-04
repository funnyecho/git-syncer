package runners

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer/contrib"
	"github.com/funnyecho/git-syncer/variable"
)

// Contrib contrib runner
func Contrib(_ []string) (runner.BubbleTap, error) {
	git, gitErr := UseGitter()
	if gitErr != nil {
		return nil, errors.WrapC(gitErr, exitcode.InvalidRunnerDependency, "failed to get gitter")
	}

	cb, cbErr := contrib.UseFactory()(git)
	if cbErr != nil {
		return nil, errors.Wrap(cbErr, "failed to init contrib")
	}

	variable.WithContrib(cb)
	return nil, nil
}

// UseContrib use remote
var UseContrib = variable.UseContrib
