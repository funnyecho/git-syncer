package syncer

import (
	"flag"
	"github.com/funnyecho/git-syncer/internal/config"
)

var remote = ""

func WithRemoteFlag(set *flag.FlagSet) {
	set.StringVar(&remote, "remote", "", "")
}

func SetupRemote() {
	config.WithRemote(remote)
}