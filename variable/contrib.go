package variable

import (
	"github.com/funnyecho/git-syncer/syncer/contrib"
)

var cb contrib.Contrib

// WithContrib set remote
func WithContrib(c contrib.Contrib) {
	cb = c
}

// UseContrib get remote
func UseContrib() (contrib.Contrib, error) {
	if cb == nil {
		return nil, ErrNotFound
	}

	return cb, nil
}
