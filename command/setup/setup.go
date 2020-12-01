package setup

import (
	"context"
	"github.com/funnyecho/git-syncer/internal/constants"
	"github.com/funnyecho/git-syncer/internal/log"
	"github.com/funnyecho/git-syncer/internal/scopex"
	"github.com/funnyecho/git-syncer/syncer"
	"github.com/mitchellh/cli"
)

func Factory() (cli.Command, error)  {
	return &cmd{}, nil
}

type cmd struct {

}

func (c *cmd) Help() string {
	return "Uploads all git-tracked non-ignored files to the remote contrib and " +
		"creates the `.git-ftp.log` file containing the SHA1 of the latest commit."
}

func (c *cmd) Synopsis() string {
	return "Setup[ remote contrib to the latest commit of repo"
}

func (c *cmd) Run(args []string) int {
	ctx := context.Background()

	if gitVerErr := syncer.CheckGitVersion(); gitVerErr != nil {
		log.Errore(ctx, gitVerErr)
		return constants.ErrorStatusGit
	}

	if projectDir, projectDirErr := syncer.CheckIsGitProject(); projectDirErr != nil {
		log.Errorw(ctx, "not a git repository (or any of the parent directories): .git", projectDirErr)
		return constants.ErrorStatusGit
	} else {
		ctx = scopex.WithProjectDir(ctx, projectDir)
	}

	return constants.ErrorStatusNoError
}
