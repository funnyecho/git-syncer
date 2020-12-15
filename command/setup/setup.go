package setup

import (
	"context"

	"github.com/funnyecho/git-syncer/internal/constants"
	"github.com/funnyecho/git-syncer/internal/log"
	"github.com/funnyecho/git-syncer/internal/scopex"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer"
	"github.com/mitchellh/cli"
)

// Factory of command `setup`
func Factory() (cli.Command, error) {
	return &cmd{}, nil
}

type cmd struct {
}

func (c *cmd) Help() string {
	return "Uploads all git-tracked non-ignored files to the remote contrib and " +
		"creates the `.git-ftp.log` file containing the SHA1 of the latest commit."
}

func (c *cmd) Synopsis() string {
	return "Setup remote contrib to the latest commit of repo"
}

func (c *cmd) Run(args []string) int {
	ctx := context.Background()

	if gitVerErr := syncer.CheckGitVersion(); gitVerErr != nil {
		log.Errore(ctx, errors.NewError(errors.WithMsg("Not a Git project? "), errors.WithErr(gitVerErr)))
		return constants.ErrorStatusGit
	}

	if projectDir, projectDirErr := syncer.CheckIsGitProject(); projectDirErr != nil {
		log.Errore(ctx, errors.NewError(errors.WithMsg("not a git repository (or any of the parent directories): .git"), errors.WithErr(projectDirErr)))
		return constants.ErrorStatusGit
	} else {
		ctx = scopex.WithProjectDir(ctx, projectDir)
	}

	if repoDirtyErr := syncer.CheckIsDirtyRepository(); repoDirtyErr != nil {
		log.Errore(ctx, errors.NewError(errors.WithMsg("Dirty repository: Having uncommitted changes. "), errors.WithErr(repoDirtyErr)))
		return constants.ErrorStatusGit
	}

	return constants.ErrorStatusNoError
}
