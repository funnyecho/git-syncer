package syncer

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer/gitter"
	"github.com/funnyecho/git-syncer/syncer/remote"
)

// Push command handler
func Push(remote remote.Remote, git gitter.Gitter) error {
	if remote == nil || git == nil {
		return errors.Err(exitcode.InvalidParams, "remote and gitter are required")
	}

	remoteSHA1, remoteSHA1Err := remote.GetHeadSHA1()
	if remoteSHA1Err != nil {
		return errors.Wrap(
			remoteSHA1Err,
			"failed to get deployed sha1",
		)
	} else if remoteSHA1 == "" {
		return errors.Err(
			exitcode.RemoteHeadNotFound,
			"deployed commit not found, use `catchup` to update deployed commit or `setup` to setup remote",
		)
	}

	headSHA1, headErr := git.GetHeadSHA1()
	if headErr != nil {
		return errors.Wrap(headErr, "failed to get repo head sha1")
	}

	uploads, deletes, filesErr := git.ListChangedFiles(getSyncRootConfig(git), remoteSHA1)
	if filesErr != nil {
		return errors.Wrap(filesErr, "failed to list changed files from commit %s", remoteSHA1)
	}

	if uploaded, deleted, syncErr := remote.Sync(headSHA1, uploads, deletes); syncErr != nil {
		return errors.Wrap(
			syncErr,
			"failed to sync changed files to remote, upload %s, delete %s",
			fmt.Sprintf("%v/%v", len(uploaded), len(uploads)),
			fmt.Sprintf("%v/%v", len(deleted), len(deletes)),
		)
	}

	return nil
}
