package alioss

import "path/filepath"

const (
	// ObjectLockFile lock file path
	ObjectLockFile = ".git-syncer/lockfile"

	// ObjectHeadLinkFile head log file path
	ObjectHeadLinkFile = ".git-syncer/head"

	// ObjectLogDir log dir path
	ObjectLogDir = ".git-syncer/logs"
)

func (a *Alioss) pathToKey(path string) string {
	return filepath.Join(a.opts.Base(), path)
}
