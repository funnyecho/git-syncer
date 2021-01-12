package gitrepo

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/fs"
)

func (r *repo) InitSyncRoot() error {
	root, _ := r.GetConfig("sync_root")
	if root == "" {
		root = "./assets"
	}

	if rootExisted, rootErr := fs.IsDirExists(root); rootExisted {
		r.syncRoot = root
		return nil
	} else {
		return errors.NewError(
			errors.WithStatusCode(exitcode.Usage),
			errors.WithErr(rootErr),
			errors.WithMsg(fmt.Sprintf("invalid sync root directory: %s", root)),
		)
	}
}
