package push

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/command/internal/runners"
	"github.com/funnyecho/git-syncer/syncer"
	"github.com/mitchellh/cli"
)

// Factory command `setup` factory
func Factory() (cli.Command, error) {
	return &cmd{}, nil
}

type cmd struct {
}

func (c *cmd) Help() string {
	return `Uploads files that have changed and
deletes files that have been deleted since the last upload.
If you are using GIT LFS, this uploads LFS link files,
not large files (stored on LFS server).
To upload the LFS tracked files, run "git lfs pull"
before "git syncer push": LFS link files will be replaced with
large files so they can be uploaded.`
}

func (c *cmd) Synopsis() string {
	return "Push changed files to remote"
}

func (c *cmd) Run(args []string) (ext int) {
	opt := &Options{}

	return runner.Run(
		args,
		runners.PrintErr,
		runners.PrintUsageErr,
		runners.WithFlagset("push", opt),
		runners.WorkingDir,
		runners.Gitter,
		runners.Remote,
		runners.WithTarget(func(_ []string) error {
			gitter, gitterErr := runners.UseGitter()
			if gitterErr != nil {
				return gitterErr
			}

			remote, remoteErr := runners.UseRemote()
			if remoteErr != nil {
				return remoteErr
			}

			return syncer.Push(remote, gitter)
		}),
	)
}
