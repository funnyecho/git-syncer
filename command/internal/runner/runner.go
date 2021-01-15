package runner

import (
	"flag"
	"fmt"

	stderrors "errors"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/flagbinder"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/repository"
	"github.com/funnyecho/git-syncer/repository/gitrepo"
)

// Run run command with taps
func Run(name string, args []string, withTaps ...WithTap) int {
	taps := defaultTaps()
	for _, wt := range withTaps {
		wt(taps)
	}

	flags := flag.NewFlagSet(name, flag.ContinueOnError)

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
		return exitcode.RepoUnknown
	}

	if options.Branch != "" {
		head, checkoutErr := repo.PushHead(options.Branch)
		if checkoutErr != nil {
			log.Errore(fmt.Sprintf("failed to checkout to branch: %s", options.Branch), checkoutErr)
			return exitcode.RepoCheckoutFailed
		}
		defer func() {
			checkoutErr := repo.PopHead(head)
			if checkoutErr != nil {
				log.Errore(fmt.Sprintf("failed to reset to head: %s", head), checkoutErr)
			}
		}()
	}

	ct := contrib.UseFactory()(repo)

	if taps.TapCommand != nil {
		cmdErr := taps.TapCommand(repo, ct)
		if cmdErr != nil {
			log.Errore("run command failed", cmdErr, "cmdName", name)

			var e *errors.Error
			stderrors.As(cmdErr, &e)
			if e == nil {
				return exitcode.Unknown
			}

			return e.Code
		}
	}

	return exitcode.Nil
}

// Taps struct to store runner taps
type Taps struct {
	TapCommand func(repository.Repository, contrib.Contrib) error
}

// WithTap wrapper to modify runner taps
type WithTap func(*Taps)

// WithTapCommand tap to command executor
func WithTapCommand(c func(repository.Repository, contrib.Contrib) error) WithTap {
	return func(t *Taps) {
		t.TapCommand = c
	}
}

func defaultTaps() *Taps {
	return &Taps{}
}
