package contrib

import "github.com/funnyecho/git-syncer/repository"

type SyncReq struct {
	SHA1    string
	Uploads []string
	Deletes []string
}

type SyncRes struct {
	SHA1     string
	Uploaded []string
	Deleted  []string
}

type Contrib interface {
	GetHeadSHA1() (string, error)
	Sync(SyncReq) (SyncRes, error)
}

type Factory func(options interface {
	repository.ConfigReader
}) Contrib

var factory Factory

func WithFactory(ft Factory) {
	factory = ft
}

func UseFactory() Factory {
	return factory
}
