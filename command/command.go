package command

import (
	"fmt"
	"github.com/funnyecho/git-syncer/constants/exitcode"
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

	c.Commands = Register()

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(exitcode.Unknown)
	} else {
		os.Exit(exitStatus)
	}
}
