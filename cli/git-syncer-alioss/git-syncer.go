package main

import (
	"github.com/funnyecho/git-syncer/command"
	"github.com/funnyecho/git-syncer/internal/contrib"
	"github.com/funnyecho/git-syncer/internal/contrib/alioss"
)

func main() {
	contrib.WithContrib(alioss.New())
	command.Run()
}
