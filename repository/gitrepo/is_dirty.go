package gitrepo

import "github.com/funnyecho/git-syncer/repository/gitrepo/gitter"

func (r *repo) IsDirtyRepository() (bool, error) {
	status, err := r.gitter.GetPorcelainStatus(gitter.WithUnoPorcelainStatus())
	if err != nil {
		return false, err
	}

	if len(status) > 0 {
		return false, nil
	}

	return true, nil
}
