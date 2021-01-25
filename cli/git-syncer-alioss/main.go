package main

import (
	"github.com/funnyecho/git-syncer/command"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/contrib/alioss"
	"github.com/funnyecho/git-syncer/pkg/log"
)

func main() {
	log.WithVerbose(log.VerboseInfo)
	contrib.WithFactory(alioss.NewContribFactory())
	command.Exec()
}
