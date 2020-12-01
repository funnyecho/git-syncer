package syncer

import (
	"github.com/funnyecho/git-syncer/internal/gitter"
)

func CheckIsGitProject() (projectDir string, err error) {
	projectDir, err = gitter.GetProjectDir()

	return
}
