package syncer

import (
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/syncer/gitter"
	"github.com/funnyecho/git-syncer/syncer/remote"
)

// Catchup command handler
func Catchup(remote remote.Remote, git gitter.Gitter) error {
	headSHA1, headErr := git.GetHeadSHA1()
	if headErr != nil {
		return headErr
	}

	if _, _, syncErr := remote.Sync(headSHA1, nil, nil); syncErr != nil {
		log.Errore("failed to catchup repo commit to remote", syncErr)
		return syncErr
	}

	return nil
}
