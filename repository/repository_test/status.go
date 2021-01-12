package repository_test

import "github.com/funnyecho/git-syncer/repository"

// NewStatus mock repo status
func NewStatus() repository.Status {
	return &mockStatus{}
}

type mockStatus struct {
}

func (m *mockStatus) IsDirtyRepository() (bool, error) {
	return false, nil
}
