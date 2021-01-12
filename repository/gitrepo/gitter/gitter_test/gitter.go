package gitter_test

import "github.com/funnyecho/git-syncer/repository/gitrepo/gitter"

// NewMockGitter new mock gitter
func NewMockGitter() gitter.Gitter {
	return &git{}
}

type git struct {
}

func (g *git) Checkout(head string) error {
	return nil
}

func (g *git) ConfigGet(withArgs ...gitter.WithArgs) (string, error) {
	return "", nil
}

func (g *git) GetHead() string {
	return ""
}

func (g *git) GetHeadSHA1() (string, error) {
	return "", nil
}

func (g *git) GetPorcelainStatus(withArgs ...gitter.WithArgs) (status []string, err error) {
	return nil, nil
}

func (g *git) GetSymbolicHead() string {
	return ""
}

func (g *git) GetVersion() (majorVersion, minorVersion int, err error) {
	return 0, 0, nil
}

func (g *git) ListFiles(path string) ([]string, error) {
	return nil, nil
}
