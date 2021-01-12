package gitter

import (
	"github.com/funnyecho/git-syncer/pkg/command"
)

func NewDefaultGitter() Gitter {
	return NewDefaultGitterWithCommander(nil)
}

func NewDefaultGitterWithCommander(cmd command.Commander) Gitter {
	if cmd == nil {
		cmd = command.UseCommand()
	}

	return &git{
		command: cmd,
	}
}

type Gitter interface {
	Checkout(head string) error
	ConfigGet(withArgs ...WithArgs) (string, error)
	ConfigSet(withArgs ...WithArgs) error
	GetHead() string
	GetHeadSHA1() (string, error)
	GetPorcelainStatus(withArgs ...WithArgs) (status []string, err error)
	GetSymbolicHead() string
	GetVersion() (majorVersion, minorVersion int, err error)

	ListFiles(path string) ([]string, error)
}

// WithArgs type for define arguments
type WithArgs func() string

type git struct {
	command command.Commander
}
