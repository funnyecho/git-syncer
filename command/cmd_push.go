package command

import (
	"flag"
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/repository"
	"github.com/mitchellh/cli"
)

// PushFactory command `push`
func PushFactory() (cli.Command, error) {
	return &pushCmd{}, nil
}

type pushCmd struct {
}

func (c *pushCmd) Help() string {
	return `Uploads files that have changed and
deletes files that have been deleted since the last upload.
If you are using GIT LFS, this uploads LFS link files,
not large files (stored on LFS server).
To upload the LFS tracked files, run "git lfs pull"
before "git ftp push": LFS link files will be replaced with
large files so they can be uploaded.`
}

func (c *pushCmd) Synopsis() string {
	return "Push changed files to remote contrib"
}

func (c *pushCmd) Run(args []string) int {
	var opt BasicOptions
	fs := flag.NewFlagSet("push", flag.ContinueOnError)

	if optCode := OptionsBinder(&opt, fs)(args); optCode != exitcode.Nil {
		return optCode
	}

	repo, repoCode := NewRepo(opt.Base, opt.Remote)
	if repoCode != exitcode.Nil {
		return repoCode
	}

	ct, ctCode := NewContrib(repo)
	if ctCode != exitcode.Nil {
		return ctCode
	}

	if setupErr := ExecPush(ct, repo); setupErr != nil {
		log.Errore("push failed", setupErr)
		return errors.GetErrorCode(setupErr)
	}

	return exitcode.Nil
}

// ExecPush push changed and deleted files to contrib
func ExecPush(c contrib.Contrib, r repository.Files) error {
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
