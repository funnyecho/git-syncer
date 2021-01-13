package setup

import (
	stderrors "errors"
	"flag"
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/flagbinder"
	"github.com/funnyecho/git-syncer/repository/gitrepo"

	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/mitchellh/cli"
)

// Factory of command `setup`
func Factory() (cli.Command, error) {
	return &cmd{}, nil
}

type cmd struct {
}

func (c *cmd) Help() string {
	return "Uploads all git-tracked non-ignored files to the remote contrib and " +
		"creates the `.git-ftp.log` file containing the SHA1 of the latest commit."
}

func (c *cmd) Synopsis() string {
	return "Setup remote contrib to the latest commit of repo"
}

func (c *cmd) Run(args []string) (ext int) {
	flags := flag.NewFlagSet("setup", flag.ContinueOnError)

	var options Options
	if bindFlagErr := flagbinder.Bind(&options, flags); bindFlagErr != nil {
		log.Errore("Failed to bind flag", bindFlagErr)
		return exitcode.Unknown
	}

	if flagErr := flags.Parse(args); flagErr != nil {
		log.Errore("Failed to parse flags", flagErr)
		return exitcode.Usage
	}

	repo, repoErr := gitrepo.New(
		gitrepo.WithWorkingDir(options.Base),
		gitrepo.WithWorkingRemote(options.Remote),
	)
	if repoErr != nil {
		log.Errore("failed setup git repo", repoErr)
		return exitcode.Git
	}

	if options.Branch != "" {
		head, checkoutErr := repo.PushHead(options.Branch)
		if checkoutErr != nil {
			log.Errore(fmt.Sprintf("failed to checkout to branch: %s", options.Branch), checkoutErr)
			return exitcode.Git
		}
		defer func() {
			checkoutErr := repo.PopHead(head)
			if checkoutErr != nil {
				log.Errore(fmt.Sprintf("failed to reset to head: %s", head), checkoutErr)
			}
		}()
	}

	ct := contrib.UseFactory()(repo)
	setupErr := contrib.Setup(ct, repo)
	if setupErr != nil {
		log.Errore("setup failed", repoErr)
		var e *errors.Error
		stderrors.As(setupErr, &e)
		if e == nil {
			return exitcode.Unknown
		} else {
			return e.StatusCode
		}
	}

	return exitcode.Nil
}
