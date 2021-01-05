package setup

import (
	"flag"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/internal/contrib"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/flagbinder"
	"github.com/funnyecho/git-syncer/repository"

	"github.com/funnyecho/git-syncer/pkg/log"
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
	repo := repository.UseRepository()
	syncer := contrib.UseContrib()

	var options Options
	if bindFlagErr := flagbinder.Bind(&options, flags); bindFlagErr != nil {
		log.Errore("Failed to bind flag", bindFlagErr)
		return exitcode.Unknown
	}

	if flagErr := flags.Parse(args); flagErr != nil {
		log.Errore("Failed to parse flags", flagErr)
		return exitcode.Usage
	}

	var projectDir string
	var currentHead string
	var remoteHeadSHA1 string
	var workingBranch string
	var workingDir string
	var workingHeadSHA1 string

	var uploadFiles []string

	// FIXME: temp to avoid `unused variables` error
	_ = projectDir
	_ = workingHeadSHA1
	_ = remoteHeadSHA1
	_ = workingBranch
	_ = workingDir

	_ = uploadFiles

	if gitMajorV, gitMinorV, gitVerErr := repo.GitVersion(); gitVerErr != nil {
		log.Errore("Git haven't installed? ", gitVerErr)
		return exitcode.Git
	} else if gitMajorV < 2 && gitMinorV < 7 {
		log.Errorw("git is too old, 1.7.0 or higher supported only")
		return exitcode.Git
	}

	if repoDir, repoDirErr := repo.GetRepoDir(); repoDirErr != nil {
		log.Errore("not a git repository (or any of the parent directories): .git", repoDirErr)
		return exitcode.Git
	} else {
		projectDir = repoDir
	}

	if repoDirty, repoDirtyErr := repo.IsDirtyRepository(); repoDirtyErr != nil || repoDirty {
		log.Errorw("Dirty repository: Having uncommitted changes. ", "err", repoDirtyErr)
		return exitcode.Git
	}

	if remoteErr := syncer.CheckAccessible(); remoteErr != nil {
		log.Errore("remote not accessible", remoteErr)
		return exitcode.RemoteForbidden
	}

	if head, headErr := repo.GetHead(); headErr != nil {
		log.Errore("Get repository head failed", headErr)
		return exitcode.Git
	} else {
		currentHead = head
	}

	if options.Branch != "" && options.Branch != currentHead {
		if setHeadErr := repo.SetHead(options.Branch); setHeadErr != nil {
			log.Errore("Set repository head failed", setHeadErr, "target head", options.Branch)
			return exitcode.Git
		} else {
			defer func() {
				r := recover()

				_ = repo.SetHead(currentHead)

				if r != nil {
					panic(r)
				}
			}()
		}
	}

	if syncRoot, syncRootErr := repo.GetSyncRoot(); syncRootErr != nil {
		log.Errorw("Sync root not a valid directory", "err", syncRootErr)
		return exitcode.Usage
	} else {
		workingDir = syncRoot
	}

	if localSHA1, localSHA1Err := repo.GetHeadSHA1(); localSHA1Err != nil {
		log.Errorw("Can't not get local revision")
		return exitcode.Git
	} else {
		workingHeadSHA1 = localSHA1
	}

	if remoteSHA1, remoteSHA1Err := syncer.GetHeadSHA1(); remoteSHA1Err != nil {
		log.Errore("Failed to check remote is clean", remoteSHA1Err)
		return exitcode.RemoteForbidden
	} else if remoteSHA1 != "" {
		log.Errorw("Commit found, use 'git syncer push' to sync")
		return exitcode.Usage
	} else {
		remoteHeadSHA1 = remoteSHA1
	}

	if lockRemoteErr := syncer.Lock(); lockRemoteErr != nil {
		log.Errore("Failed to lock remote", lockRemoteErr)
		return exitcode.RemoteForbidden
	} else {
		defer func() {
			r := recover()

			unlockErr := syncer.Unlock()
			if unlockErr != nil {
				panic(errors.NewError(
					errors.WithMsg("Failed to unlock remote"),
					errors.WithErr(unlockErr),
				))
			}

			if r != nil {
				panic(r)
			}
		}()
	}

	if allFiles, listFilesErr := repo.ListAllFiles(projectDir); listFilesErr != nil {
		log.Errore("Failed to list files", listFilesErr)
		return exitcode.Git
	} else {
		uploadFiles = allFiles
	}

	if uploadFilesErr := syncer.UploadFiles(uploadFiles); uploadFilesErr != nil {
		// FIXME: try to rollback partial uploaded files
		log.Errore("Failed to upload files", uploadFilesErr)
		return exitcode.Upload
	}

	if setRemoteSHA1Err := syncer.SetHeadSHA1(workingHeadSHA1); setRemoteSHA1Err != nil {
		// FIXME: maybe try to rollback uploaded files
		log.Errore("Failed to set remote sha1", setRemoteSHA1Err)
		return exitcode.Upload
	}

	return exitcode.Nil
}
