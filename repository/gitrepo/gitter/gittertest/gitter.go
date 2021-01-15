package gittertest

import "github.com/funnyecho/git-syncer/repository/gitrepo/gitter"

// New new mock gitter
func New() gitter.Gitter {
	return &git{}
}

type git struct {
}

func (g *git) Checkout(head string) error {
	return nil
}

func (g *git) ConfigGet(key string, options gitter.ConfigGetOptions) (string, error) {
	return "", nil
}

func (g *git) ConfigSet(key, val string, options gitter.ConfigSetOptions) error {
	return nil
}

func (g *git) GetHead() (string, error) {
	return "", nil
}

func (g *git) GetHeadSHA1() (string, error) {
	return "", nil
}

func (g *git) GetPorcelainStatus() (status []string, err error) {
	return nil, nil
}

func (g *git) GetUnoPorcelainStatus() (status []string, err error) {
	return nil, nil
}

func (g *git) GetSymbolicHead() (string, error) {
	return "", nil
}

func (g *git) GetVersion() (majorVersion, minorVersion int, err error) {
	return 0, 0, nil
}

func (g *git) ListFiles(path string) ([]string, error) {
	return nil, nil
}

func (g *git) DiffAM(path string, commit string) ([]string, error) {
	return nil, nil
}

func (g *git) DiffD(path string, commit string) ([]string, error) {
	return nil, nil
}
