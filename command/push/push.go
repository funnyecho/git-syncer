package push

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/repository"
	"github.com/mitchellh/cli"
)

// Factory of command `push`
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
before "git ftp push": LFS link files will be replaced with 
large files so they can be uploaded.`
}

func (c *cmd) Synopsis() string {
	return "Push changed files to remote contrib"
}

func (c *cmd) Run(args []string) int {
	return runner.Run("push", args, runner.WithTapCommand(func(r repository.Repository, c contrib.Contrib) error {
		return contrib.Push(c, r)
	}))
}
