package command

import (
	"fmt"
	"github.com/funnyecho/git-syncer/constants"
	"github.com/funnyecho/git-syncer/repository"
	"github.com/funnyecho/git-syncer/repository/repo"
	"github.com/mitchellh/cli"
	"os"
)

var (
	BuildTime     string
	Version       string
	BuildPlatform string
)

func Run() {
	c := cli.NewCLI("git-syncer", Version)
	c.Args = os.Args[1:]

	repository.PushRepository(gitrepo.New())
	defer func() {
		repository.PopRepository(nil)
	}()

	c.Commands = Register()

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(constants.ErrorStatusUnknown)
	} else {
		os.Exit(exitStatus)
	}
}
