package syncer

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/syncer/gitter"
	"github.com/funnyecho/git-syncer/syncer/remote"
)

// Setup command handler
func Setup(remote remote.Remote, git gitter.Gitter) error {
	if sha1, sha1Err := remote.GetHeadSHA1(); sha1Err != nil {
		return errors.WrapC(
			sha1Err,
			exitcode.RemoteForbidden,
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
		return headErr
	}

	uploads, filesErr := git.ListTrackedFiles(getSyncRootConfig(git))
	if filesErr != nil {
		return filesErr
	}

	if uploaded, _, syncErr := remote.Sync(headSHA1, uploads, nil); syncErr != nil {
		log.Errore(
			"failed to sync tracked files to remote",
			syncErr,
			"upload", fmt.Sprintf("%v/%v", len(uploaded), len(uploads)),
		)
		return syncErr
	}

	return nil
}
