package runner

import (
	"flag"
	"fmt"

	"github.com/funnyecho/git-syncer/command/internal/flagparser"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/pkg/errors"
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

	if argParseErr := flagparser.Parser(&options, flags)(args); argParseErr != nil {
		log.Errore("Failed to parse flags", argParseErr)
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

	ct, ctErr := contrib.UseFactory()(repo)
	if ctErr != nil {
		log.Errore(fmt.Sprintf("failed to init contrib"), ctErr)
		return exitcode.ContribUnknown
	}

	if taps.TapCommand != nil {
		cmdErr := taps.TapCommand(repo, ct)
		if cmdErr != nil {
			log.Errore("run command failed", cmdErr, "cmdName", name)

			return errors.GetErrorCode(cmdErr)
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
