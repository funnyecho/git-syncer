package gitrepo

import (
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
)

// IsDirtyRepository check whether repository is dirty
func IsDirtyRepository(git gitter.Status) (bool, error) {
	status, err := git.GetUnoPorcelainStatus()
	if err != nil {
		return false, err
	}

	if len(status) > 0 {
		return true, nil
	}

	return false, nil
}
