package syncer

import (
	"fmt"

	"github.com/funnyecho/git-syncer/internal/constants"
	"github.com/funnyecho/git-syncer/internal/errors"
	"github.com/funnyecho/git-syncer/internal/gitter"
)

// CheckIsDirtyRepository check whether has modified files
func CheckIsDirtyRepository() error {
	status, err := gitter.GetPorcelainStatus(gitter.WithUnoPorcelainStatus())
	if err != nil {
		return err
	}

	if len(status) > 0 {
		return errors.NewError(
			errors.WithStatusCode(constants.ErrorStatusGit),
			errors.WithMsg(fmt.Sprintf("uno status: %v+", status)),
		)
	}

	return nil
}
