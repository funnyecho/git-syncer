package catchup

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/repository"
)

// Catchup catchup contrib with repo head
func Catchup(c contrib.Contrib, r repository.HeadReader) error {
	sha1, sha1Err := r.GetHeadSHA1()

	if sha1Err != nil {
		return errors.NewError(
			errors.WithMsg("failed to get repo sha1"),
			errors.WithErr(sha1Err),
		)
	} else if sha1 == "" {
		return errors.NewError(
			errors.WithCode(exitcode.RepoHeadNotFound),
			errors.WithMsg("repo sha1 can't be empty"),
		)
	}

	res, syncErr := c.Sync(&contrib.SyncReq{
		SHA1:    sha1,
		Uploads: nil,
		Deletes: nil,
	})
	if syncErr != nil {
		return errors.NewError(
			errors.WithCode(exitcode.ContribSyncFailed),
			errors.WithMsg(fmt.Sprintln(
				"failed to sync sha1 of repo. try catchup later",
				fmt.Sprintf("deployedSHA1: %s", res.SHA1),
			)),
		)
	}

	return nil
}
