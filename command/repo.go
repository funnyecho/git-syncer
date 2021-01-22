package command

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/repository"
	"github.com/funnyecho/git-syncer/repository/gitrepo"
)

// NewRepo new repo instance
func NewRepo(base, remote string) (repository.Repository, int) {
	repo, repoErr := gitrepo.New(
		gitrepo.WithWorkingDir(base),
		gitrepo.WithWorkingRemote(remote),
	)

	if repoErr != nil {
		log.Errore("failed setup git repo", repoErr)
		return nil, exitcode.RepoUnknown
	}

	return repo, exitcode.Nil
}
