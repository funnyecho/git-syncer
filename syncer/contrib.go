package syncer

import "github.com/funnyecho/git-syncer/internal/contrib"

var wipContrib contrib.Contrib

func SetupContrib(c contrib.Contrib) {
	wipContrib = c
}
