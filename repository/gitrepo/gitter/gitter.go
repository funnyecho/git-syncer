package gitter

import (
	"github.com/funnyecho/git-syncer/pkg/command"
)

func NewDefaultGitter() Gitter {
	return &git{
		command: command.Command,
	}
}

type Gitter interface {
	GetVersion() (majorVersion, minorVersion int, err error)
	ListFiles(path string) ([]string, error)
}

// WithArgs type for define arguments
type WithArgs func() string

type git struct {
	command command.Commander
}
