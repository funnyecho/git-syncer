package main

import (
	"fmt"
	"github.com/funnyecho/git-syncer/command"
	"github.com/funnyecho/git-syncer/constants"
	"github.com/mitchellh/cli"
	"os"
)

var (
	BuildTime     string
	Version       string
	BuildPlatform string
)

func main() {
	c := cli.NewCLI("git-syncer", Version)
	c.Args = os.Args[1:]

	c.Commands = command.Register()

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(constants.ErrorStatusUnknown)
	} else {
		os.Exit(exitStatus)
	}
}
