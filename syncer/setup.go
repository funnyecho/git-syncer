package syncer

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer/contrib"
	"github.com/funnyecho/git-syncer/syncer/gitter"
)

// Setup command handler
func Setup(cb contrib.Contrib, git gitter.Gitter) error {
	if cb == nil || git == nil {
		return errors.Err(exitcode.InvalidParams, "contrib and gitter are required")
	}

	if sha1, sha1Err := cb.GetHeadSHA1(); sha1Err != nil {
		return errors.Wrap(
			sha1Err,
			"failed to get deployed sha1",
		)
	} else if sha1 != "" {
		return errors.Err(
			exitcode.Usage,
			"deployed commit found, use 'push' to sync",
		)
	}

	headSHA1, headErr := git.GetHeadSHA1()
	if headErr != nil {
		return errors.Wrap(headErr, "failed to get repo head sha1")
	}

	uploads, filesErr := git.ListTrackedFiles(getSyncRootConfig(git))
	if filesErr != nil {
		return errors.Wrap(filesErr, "failed to list tracked files from repo")
	}

	if uploaded, _, syncErr := cb.Sync(headSHA1, uploads, nil); syncErr != nil {
		return errors.Wrap(syncErr, "failed to sync tracked files to contrib, upload %s ", fmt.Sprintf("%v/%v", len(uploaded), len(uploads)))
	}

	return nil
}
