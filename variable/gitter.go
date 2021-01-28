package variable

import "github.com/funnyecho/git-syncer/syncer/gitter"

var gt gitter.Gitter

// WithGitter set gitter
func WithGitter(git gitter.Gitter) {
	gt = git
}

// UseGitter get gitter
func UseGitter() (gitter.Gitter, error) {
	if gt == nil {
		return nil, ErrNotFound
	}

	return gt, nil
}
