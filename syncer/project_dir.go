package syncer

import (
	"github.com/funnyecho/git-syncer/pkg/gitter"
)

var projectDir = ""

func SetupProjectDir() error {
	d, err := gitter.GetProjectDir()
	if err != nil {
		return err
	} else {
		projectDir = d
	}

	return nil
}
