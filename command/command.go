package command

import (
	"fmt"
	"os"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/mitchellh/cli"
)

var (
	BuildTime     string
	Version       string
	BuildPlatform string
)

// Exec exec command
func Exec() {
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
