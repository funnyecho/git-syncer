package command

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/repository"
)

func NewContrib(repo repository.Repository) (contrib.Contrib, int) {
	ct, ctErr := contrib.UseFactory()(repo)
	if ctErr != nil {
		log.Errore(fmt.Sprintf("failed to init contrib"), ctErr)
		return nil, exitcode.ContribForbidden
	}

	return ct, exitcode.Nil
}
