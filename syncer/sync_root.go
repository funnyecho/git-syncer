package syncer

import (
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/fs"
)

var syncRoot = ""

func SetSyncRoot() error {
	r, _ := GetConfig("sync_root")
	if r == "" {
		r = "./assets"
	}

	if rootExisted, rootErr := fs.IsDirExists(r); rootExisted {
		syncRoot = r
		return nil
	} else {
		return errors.NewError(errors.WithMsg("sync root not an valid directory"), errors.WithErr(rootErr))
	}
}