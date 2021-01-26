package syncer

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/syncer/gitter"
	"github.com/funnyecho/git-syncer/syncer/remote"
)

// Push command handler
func Push(remote remote.Remote, git gitter.Gitter) error {
	remoteSHA1, remoteSHA1Err := remote.GetHeadSHA1()
	if remoteSHA1Err != nil {
		return errors.WrapC(
			remoteSHA1Err,
			exitcode.RemoteForbidden,
			"failed to get deployed sha1",
		)
	} else if remoteSHA1 == "" {
		return errors.Err(
			exitcode.Usage,
			"deployed commit not found, use `catchup` to update deployed commit or `setup` to setup remote",
		)
	}

	headSHA1, headErr := git.GetHeadSHA1()
	if headErr != nil {
		return headErr
	}

	uploads, deletes, filesErr := git.ListChangedFiles(getSyncRootConfig(git), remoteSHA1)
	if filesErr != nil {
		return filesErr
	}

	if uploaded, deleted, syncErr := remote.Sync(headSHA1, uploads, deletes); syncErr != nil {
		log.Errore(
			"failed to sync tracked files to remote",
			syncErr,
			"upload", fmt.Sprintf("%v/%v", len(uploaded), len(uploads)),
			"delete", fmt.Sprintf("%v/%v", len(deleted), len(deletes)),
		)
		return syncErr
	}

	return nil
}
