package main

import (
	"github.com/funnyecho/git-syncer/command"
	"github.com/funnyecho/git-syncer/internal/contrib"
	"github.com/funnyecho/git-syncer/internal/contrib/alioss"
	"github.com/funnyecho/git-syncer/repository"
	"github.com/funnyecho/git-syncer/repository/repo"
)

func main() {
	contrib.WithContrib(alioss.New())
	repository.WithRepository(repo.New())
	command.Run()
}
