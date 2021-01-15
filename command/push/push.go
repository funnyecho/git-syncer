package push

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/repository"
)

// Push push changed and deleted files to contrib
func Push(c contrib.Contrib, r repository.Files) error {
	contribHead, contribHeadErr := c.GetHeadSHA1()
	if contribHeadErr != nil {
		return errors.NewError(
			errors.WithMsg("failed to get contrib head sha1"),
			errors.WithErr(contribHeadErr),
		)
	} else if contribHead == "" {
		return errors.NewError(
			errors.WithMsg("contrib head is empty, try `setup` command instead"),
			errors.WithCode(exitcode.ContribHeadNotFound),
		)
	}

	repoSHA1, uploads, deletes, diffErr := r.ListChangedFiles(contribHead)
	if diffErr != nil {
		return errors.NewError(
			errors.WithMsg(fmt.Sprintf("failed to diff files from head contrib:%s to repo", contribHead)),
			errors.WithErr(diffErr),
		)
	}

	res, syncErr := c.Sync(&contrib.SyncReq{
		SHA1:    repoSHA1,
		Uploads: uploads,
		Deletes: deletes,
	})

	if syncErr != nil {
		return errors.NewError(
			errors.WithCode(exitcode.ContribSyncFailed),
			errors.WithMsg(fmt.Sprintln(
				"failed to sync changed files of repo. try later",
				fmt.Sprintf("deployedSHA1: %s", res.SHA1),
				fmt.Sprintf("uploaded files: %v", res.Uploaded),
				fmt.Sprintf("deleted files: %v", res.Deleted),
			)),
		)
	}

	return nil
}
