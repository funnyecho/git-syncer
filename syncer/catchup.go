package syncer

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/syncer/contrib"
	"github.com/funnyecho/git-syncer/syncer/gitter"
)

// Catchup command handler
func Catchup(cb contrib.Contrib, git gitter.Gitter) error {
	if cb == nil || git == nil {
		return errors.Err(exitcode.InvalidParams, "contrib and gitter are required")
	}

	headSHA1, headErr := git.GetHeadSHA1()
	if headErr != nil {
		return errors.Wrap(headErr, "failed to get repo head sha1")
	}

	if _, _, syncErr := cb.Sync(headSHA1, nil, nil); syncErr != nil {
		log.Errore("failed to catchup repo commit to remote", syncErr)
		return syncErr
	}

	return nil
}
