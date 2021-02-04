package syncer

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer/contrib"
	"github.com/funnyecho/git-syncer/syncer/gitter"
)

// Push command handler
func Push(cb contrib.Contrib, git gitter.Gitter) error {
	if cb == nil || git == nil {
		return errors.Err(exitcode.InvalidParams, "contrib and gitter are required")
	}

	contribSHA1, contribSHA1Err := cb.GetHeadSHA1()
	if contribSHA1Err != nil {
		return errors.Wrap(
			contribSHA1Err,
			"failed to get deployed sha1",
		)
	} else if contribSHA1 == "" {
		return errors.Err(
			exitcode.ContribHeadNotFound,
			"deployed commit not found, use `catchup` to update deployed commit or `setup` to setup contrib",
		)
	}

	headSHA1, headErr := git.GetHeadSHA1()
	if headErr != nil {
		return errors.Wrap(headErr, "failed to get repo head sha1")
	}

	uploads, deletes, filesErr := git.ListChangedFiles(getSyncRootConfig(git), contribSHA1)
	if filesErr != nil {
		return errors.Wrap(filesErr, "failed to list changed files from commit %s", contribSHA1)
	}

	if uploaded, deleted, syncErr := cb.Sync(headSHA1, uploads, deletes); syncErr != nil {
		return errors.Wrap(
			syncErr,
			"failed to sync changed files to contrib, upload %s, delete %s",
			fmt.Sprintf("%v/%v", len(uploaded), len(uploads)),
			fmt.Sprintf("%v/%v", len(deleted), len(deletes)),
		)
	}

	return nil
}
