package syncer

import "flag"

var remote = ""

func WithRemoteFlag(set *flag.FlagSet) {
	set.StringVar(&remote, "remote", "", "")
}