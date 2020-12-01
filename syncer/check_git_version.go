package syncer

import (
	"fmt"
	"github.com/funnyecho/git-syncer/internal/gitter"
)

func CheckGitVersion() error {
	majorV, minorV, err := gitter.GetGitVersion()

	if err != nil {
		return err
	}

	if majorV < 2 && minorV < 7 {
		return fmt.Errorf("git is too old, 1.7.0 or higher supported only")
	}

	return nil
}
