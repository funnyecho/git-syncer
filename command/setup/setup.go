package setup

import (
	"flag"

	"github.com/funnyecho/git-syncer/constants"
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
	syncer.WithBranchFlag(flags)

	if flagErr := flags.Parse(args); flagErr != nil {
		log.Errorw("Failed to parse flags", "err", flagErr)
		return constants.ErrorStatusUsage
	}

	if gitVerErr := syncer.CheckGitVersion(); gitVerErr != nil {
		log.Errorw("Not a Git project? ", "err", gitVerErr)
		return constants.ErrorStatusGit
	}

	if projectDirErr := syncer.SetupProjectDir(); projectDirErr != nil {
		log.Errorw("not a git repository (or any of the parent directories): .git", "err", projectDirErr)
		return constants.ErrorStatusGit
	}

	if repoDirtyErr := syncer.CheckIsDirtyRepository(); repoDirtyErr != nil {
		log.Errorw("Dirty repository: Having uncommitted changes. ", "err", repoDirtyErr)
		return constants.ErrorStatusGit
	}

	if syncRootErr := syncer.SetSyncRoot(); syncRootErr != nil {
		log.Errorw("Sync root not a valid directory", "err", syncRootErr)
		return constants.ErrorStatusUsage
	}

	if branchErr := syncer.SetWorkingBranch(); branchErr != nil {
		log.Errorw("Setup working branch failed", "err", branchErr)
		return constants.ErrorStatusGit
	}

	if localSHA1Err := syncer.SetupLocalSHA1(); localSHA1Err != nil {
		log.Errorw("Can't not get local revision")
		return constants.ErrorStatusGit
	}

	if contribErr := syncer.SetupContrib(); contribErr != nil {
		log.Errorw("Can't not setup contrib")
		return constants.ErrorStatusUsage
	}

	return constants.ErrorStatusNoError
}
