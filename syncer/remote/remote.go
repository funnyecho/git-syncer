package remote

import "github.com/funnyecho/git-syncer/syncer/gitter"

// Remote interface to syncer remote
type Remote interface {
	GetHeadSHA1() (string, error)
	Sync(sha1 string, uploads []string, deletes []string) (uploaded []string, deleted []string, err error)
}

// Factory factory function to new contrib instance
type Factory func(options interface {
	gitter.ConfigReader
}) (Remote, error)

var factory Factory

// WithFactory set factory singleton
func WithFactory(ft Factory) {
	factory = ft
}

// UseFactory get factory singleton
func UseFactory() Factory {
	return factory
}
