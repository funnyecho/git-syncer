package gitrepo

// DefaultSyncRoot default sync root dir
const DefaultSyncRoot = "./assets"

func (r *repo) GetSyncRoot() string {
	root, _ := r.GetConfig(ConfigSyncRoot)
	if root == "" {
		root = DefaultSyncRoot
	}

	return root
}
