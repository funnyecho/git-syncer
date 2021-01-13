package gitrepo

func (r *repo) IsDirtyRepository() (bool, error) {
	status, err := r.gitter.GetUnoPorcelainStatus()
	if err != nil {
		return false, err
	}

	if len(status) > 0 {
		return true, nil
	}

	return false, nil
}
