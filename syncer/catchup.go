package syncer

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/syncer/gitter"
	"github.com/funnyecho/git-syncer/syncer/remote"
)

// Catchup command handler
func Catchup(remote remote.Remote, git gitter.Gitter) error {
	if remote == nil || git == nil {
		return errors.Err(exitcode.InvalidParams, "remote and gitter are required")
	}

	headSHA1, headErr := git.GetHeadSHA1()
	if headErr != nil {
		return errors.Wrap(headErr, "failed to get repo head sha1")
	}

	if _, _, syncErr := remote.Sync(headSHA1, nil, nil); syncErr != nil {
		log.Errore("failed to catchup repo commit to remote", syncErr)
		return syncErr
	}

	return nil
}
