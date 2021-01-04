package alioss

import contrib2 "github.com/funnyecho/git-syncer/internal/contrib"

func New() contrib2.Contrib {
	return &contrib{}
}

type contrib struct {
}

func (c contrib) CheckAccess() error {
	panic("implement me")
}
