package gitrepo

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
)

// IsDirtyRepository check whether repository is dirty
func IsDirtyRepository(git gitter.Status) (bool, error) {
	status, err := git.GetUnoPorcelainStatus()
	if err != nil {
		return false, errors.NewError(errors.WithCode(exitcode.RepoUnknown), errors.WithErr(err))
	}

	if len(status) > 0 {
		return true, nil
	}

	return false, nil
}
