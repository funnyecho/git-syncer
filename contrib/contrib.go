package contrib

import "github.com/funnyecho/git-syncer/repository"

// SyncReq sync request info
type SyncReq struct {
	SHA1    string
	Uploads []string
	Deletes []string
}

// SyncRes sync response info
type SyncRes struct {
	SHA1     string
	Uploaded []string
	Deleted  []string
}

// Contrib contrib to be implement
type Contrib interface {
	GetHeadSHA1() (string, error)
	Sync(*SyncReq) (SyncRes, error)
}

// Factory factory function to new contrib instance
type Factory func(options interface {
	repository.ConfigReader
}) Contrib

var factory Factory

// WithFactory set factory singleton
func WithFactory(ft Factory) {
	factory = ft
}

// UseFactory get factory singleton
func UseFactory() Factory {
	return factory
}
