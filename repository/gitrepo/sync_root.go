package gitrepo

import "github.com/funnyecho/git-syncer/repository"

// DefaultSyncRoot default sync root dir
const DefaultSyncRoot = "./assets"

// GetSyncRoot return sync root
func GetSyncRoot(r repository.ConfigReader) string {
	root, _ := r.GetConfig(ConfigSyncRoot)
	if root == "" {
		root = DefaultSyncRoot
	}

	return root
}
