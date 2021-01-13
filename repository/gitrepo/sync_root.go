package gitrepo

func (r *repo) GetSyncRoot() string {
	root, _ := r.GetConfig("sync_root")
	if root == "" {
		root = "./assets"
	}

	return root
}
