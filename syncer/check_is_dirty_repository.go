package syncer

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/gitter"
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
