package repository_test

import "github.com/funnyecho/git-syncer/repository"

// NewFiles mock repo files
func NewFiles() repository.Files {
	return &mockFiles{}
}

type mockFiles struct{}

func (m *mockFiles) ListAllFiles() (sha1 string, uploads []string, err error) {
	return "", nil, nil
}

func (m *mockFiles) ListChangedFiles(baseSha1 string) (sha1 string, uploads []string, deletes []string, err error) {
	return "", nil, nil, nil
}
