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

// CatchupFactory of command `catchup`
func CatchupFactory() (cli.Command, error) {
	return &catchupCmd{}, nil
}

type catchupCmd struct {
}

func (c *catchupCmd) Help() string {
	return "Creates or updates the `.git-ftp.log` file on the remote contrib.\n" +
		"It assumes that you uploaded all other files already.\n" +
		"You might have done that with another program."
}

func (c *catchupCmd) Synopsis() string {
	return "Creates or updates the `.git-ftp.log` file on the remote contrib."
}

func (c *catchupCmd) Run(args []string) int {
	var opt BasicOptions
	fs := flag.NewFlagSet("catchup", flag.ContinueOnError)

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

	if setupErr := ExecCatchup(ct, repo); setupErr != nil {
		log.Errore("push failed", setupErr)
		return errors.GetErrorCode(setupErr)
	}

	return exitcode.Nil
}

// ExecCatchup catchup contrib with repo head
func ExecCatchup(c contrib.Contrib, r repository.HeadReader) error {
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
