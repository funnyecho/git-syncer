package alioss

import "path/filepath"

const (
	objectLockFile     = ".git-syncer/lockfile"
	objectHeadLinkFile = ".git-syncer/head"
	objectLogDir       = ".git-syncer/logs"
)

func (a *Alioss) pathToKey(path string) string {
	return filepath.Join(a.Options.Base, path)
}
