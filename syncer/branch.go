package syncer

import (
	"flag"
	"github.com/funnyecho/git-syncer/pkg/gitter"
)

var currentHead = ""
var wipBranch = ""

func WithBranchFlag(set *flag.FlagSet) {
	set.StringVar(&wipBranch, "branch", "", "")
}

func SetWorkingBranch() error {
	if wipBranch == "" {
		wipBranch, _ = GetConfig("branch")
	}

	if wipBranch != "" {
		if symbolicHead := gitter.GetSymbolicHead(); symbolicHead != "" {
			currentHead = symbolicHead
		} else if headRev := gitter.GetHead(); headRev != "" {
			currentHead = headRev
		} else if localSHA1, localSHA1Err := getLocalSHA1(); localSHA1Err == nil {
			currentHead = localSHA1
		}

		return gitter.Checkout(wipBranch)
	}

	return nil
}

func ResetWorkingBranch() error {
	if wipBranch != "" && currentHead != "" {
		return gitter.Checkout(currentHead)
	}

	return nil
}
