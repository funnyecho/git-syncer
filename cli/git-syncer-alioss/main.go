package main

import (
	"github.com/funnyecho/git-syncer/command"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/contrib/alioss"
)

func main() {
	contrib.WithFactory(alioss.NewContribFactory())
	command.Exec()
}
