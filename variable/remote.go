package variable

import "github.com/funnyecho/git-syncer/syncer/remote"

var rm remote.Remote

// WithRemote set remote
func WithRemote(r remote.Remote) {
	rm = r
}

// UseRemote get remote
func UseRemote() (remote.Remote, error) {
	if rm == nil {
		return nil, ErrNotFound
	}

	return rm, nil
}
