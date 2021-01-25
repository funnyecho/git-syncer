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

// SetupFactory command `setup` factory
func SetupFactory() (cli.Command, error) {
	return &setupCmd{}, nil
}

type setupCmd struct {
}

func (c *setupCmd) Help() string {
	return "Uploads all git-tracked non-ignored files to the remote contrib and " +
		"creates the `.git-syncer.log` file containing the SHA1 of the latest commit."
}

func (c *setupCmd) Synopsis() string {
	return "Setup remote contrib to the latest commit of repo"
}

func (c *setupCmd) Run(args []string) (ext int) {
	var opt BasicOptions
	fs := flag.NewFlagSet("setup", flag.ContinueOnError)

	if optCode := OptionsBinder(&opt, fs)(args); optCode != exitcode.Nil {
		return optCode
	}

	repo, repoCode := NewRepo(opt.Base, opt.Remote)
	if repoCode != exitcode.Nil {
		return repoCode
	}

	if opt.Branch != "" {
		prevHead, pushHeadErr := repo.PushHead(opt.Branch)
		if pushHeadErr != nil {
			log.Errore("failed to checkout to branch", pushHeadErr, "branch", opt.Branch)
			return exitcode.RepoCheckoutFailed
		}
		defer func() {
			popHeadErr := repo.PopHead(prevHead)
			if popHeadErr != nil {
				log.Errore("failed to reset to head", popHeadErr, "head", prevHead)
				if ext != exitcode.Nil {
					ext = exitcode.RepoCheckoutFailed
				}
			}
		}()
	}

	ct, ctCode := NewContrib(repo)
	if ctCode != exitcode.Nil {
		return ctCode
	}

	if setupErr := ExecSetup(ct, repo); setupErr != nil {
		log.Errore("setup failed", setupErr)
		return errors.GetErrorCode(setupErr)
	}

	return exitcode.Nil
}

// ExecSetup setup command handler
func ExecSetup(c contrib.Contrib, repo repository.Files) error {
	if sha1, sha1Err := c.GetHeadSHA1(); sha1Err != nil {
		return errors.NewError(
			errors.WithMsg("failed to get deployed sha1"),
			errors.WithErr(sha1Err),
		)
	} else if sha1 != "" {
		return errors.NewError(
			errors.WithCode(exitcode.Usage),
			errors.WithMsg("deployed commit found, use 'push' to sync."),
		)
	}

	repoSha1, files, filesErr := repo.ListAllFiles()
	if filesErr != nil {
		return filesErr
	}

	res, syncErr := c.Sync(&contrib.SyncReq{
		SHA1:    repoSha1,
		Uploads: files,
		Deletes: nil,
	})
	if syncErr != nil {
		return errors.NewError(
			errors.WithCode(exitcode.ContribSyncFailed),
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
