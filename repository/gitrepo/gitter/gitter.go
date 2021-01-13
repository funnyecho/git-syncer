package gitter

import (
	"github.com/funnyecho/git-syncer/pkg/command"
)

// NewDefaultGitter new default gitter with default commander
func NewDefaultGitter() Gitter {
	return NewDefaultGitterWithCommander(nil)
}

// NewDefaultGitterWithCommander new default gitter with specific commander
func NewDefaultGitterWithCommander(cmd command.Commander) Gitter {
	if cmd == nil {
		cmd = command.UseCommand()
	}

	return &git{
		command: cmd,
	}
}

// Gitter interface to interact with git repo
type Gitter interface {
	Config
	Status

	Checkout(head string) error
	GetHead() string
	GetHeadSHA1() (string, error)
	GetSymbolicHead() string
	GetVersion() (majorVersion, minorVersion int, err error)

	ListFiles(path string) ([]string, error)
}

type git struct {
	command command.Commander
}
