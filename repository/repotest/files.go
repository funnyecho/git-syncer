package repotest

import "github.com/funnyecho/git-syncer/repository"

// NewFiles mock repo files
func NewFiles() repository.Files {
	return &MockFiles{}
}

// MockFiles mock files
type MockFiles struct{}

// ListAllFiles list all fiels
func (m *MockFiles) ListAllFiles() (sha1 string, uploads []string, err error) {
	return "", nil, nil
}

// ListChangedFiles list changed files
func (m *MockFiles) ListChangedFiles(baseSha1 string) (sha1 string, uploads []string, deletes []string, err error) {
	return "", nil, nil, nil
}
