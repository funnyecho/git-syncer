package syncer

import (
	"github.com/funnyecho/git-syncer/pkg/gitter"
)

func CheckIsGitProject() (projectDir string, err error) {
	projectDir, err = gitter.GetProjectDir()

	return
}
