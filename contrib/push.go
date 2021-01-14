package contrib

import (
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/repository"
)

// Push push changed and deleted files to contrib
func Push(c Contrib, r repository.Files) error {
	contribHead, contribHeadErr := c.GetHeadSHA1()
	if contribHeadErr != nil {
		return errors.NewError(
			errors.WithMsg("failed to get contrib head sha1"),
			errors.WithErr(contribHeadErr),
		)
	}

	// Fixme: `contribHead` to be used
	_ = contribHead

	panic("please implement me")
}
