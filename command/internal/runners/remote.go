package runners

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer/remote"
)

var rm remote.Remote

// Remote remote runner
func Remote(_ []string) (runner.BubbleTap, error) {
	git, gitErr := UseGitter()
	if gitErr != nil {
		return nil, errors.WrapC(gitErr, exitcode.InvalidRunnerDependency, "failed to get gitter")
	}

	r, rErr := remote.UseFactory()(git)
	if rErr != nil {
		return nil, errors.Wrap(rErr, "failed to init remote")
	}

	rm = r
	return nil, nil
}

// UseRemote use remote
func UseRemote() (remote.Remote, error) {
	if rm == nil {
		return nil, errors.Err(exitcode.InvalidRunnerDependency, "remote not init")
	}

	return rm, nil
}
