package contrib

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/repository"
)

// Setup setup command handler
func Setup(c Contrib, repo repository.Files) error {
	if sha1, sha1Err := c.GetHeadSHA1(); sha1Err != nil {
		return errors.NewError(
			errors.WithStatusCode(exitcode.RemoteForbidden),
			errors.WithMsg("failed to get deployed sha1"),
			errors.WithErr(sha1Err),
		)
	} else if sha1 != "" {
		return errors.NewError(
			errors.WithStatusCode(exitcode.Usage),
			errors.WithMsg("deployed commit found, use 'push' to sync."),
		)
	}

	repoSha1, files, filesErr := repo.ListAllFiles()
	if filesErr != nil {
		return filesErr
	}

	res, syncErr := c.Sync(&SyncReq{
		SHA1:    repoSha1,
		Uploads: files,
		Deletes: nil,
	})
	if syncErr != nil {
		return errors.NewError(
			errors.WithStatusCode(exitcode.Upload),
			errors.WithMsg(fmt.Sprintln(
				"failed to sync all files of repo. try setup later",
				fmt.Sprintf("deployedSHA1: %s", res.SHA1),
				fmt.Sprintf("uploaded files: %v", res.Uploaded),
				fmt.Sprintf("deleted files: %v", res.Deleted),
			)),
		)
	}

	return nil
}
