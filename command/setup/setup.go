package setup

import (
	"flag"

	"github.com/funnyecho/git-syncer/internal/constants"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
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
	flags := flag.NewFlagSet("setup", flag.ContinueOnError)

	syncer.WithRemoteFlag(flags)

	if flagErr := flags.Parse(args); flagErr != nil {
		log.Errore(errors.NewError(errors.WithMsg("Failed to parse flags"), errors.WithErr(flagErr)))
		return constants.ErrorStatusUsage
	}

	if gitVerErr := syncer.CheckGitVersion(); gitVerErr != nil {
		log.Errore(errors.NewError(errors.WithMsg("Not a Git project? "), errors.WithErr(gitVerErr)))
		return constants.ErrorStatusGit
	}

	if projectDirErr := syncer.SetupProjectDir(); projectDirErr != nil {
		log.Errore(errors.NewError(errors.WithMsg("not a git repository (or any of the parent directories): .git"), errors.WithErr(projectDirErr)))
		return constants.ErrorStatusGit
	}

	if repoDirtyErr := syncer.CheckIsDirtyRepository(); repoDirtyErr != nil {
		log.Errore(errors.NewError(errors.WithMsg("Dirty repository: Having uncommitted changes. "), errors.WithErr(repoDirtyErr)))
		return constants.ErrorStatusGit
	}

	if syncRootErr := syncer.SetupSyncRoot(); syncRootErr != nil {
		log.Errore(errors.NewError(errors.WithMsg("Sync root not a valid directory"), errors.WithErr(syncRootErr)))
		return constants.ErrorStatusUsage
	}


	return constants.ErrorStatusNoError
}
