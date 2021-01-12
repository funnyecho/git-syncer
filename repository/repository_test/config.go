package repository_test

import (
	"fmt"

	"github.com/funnyecho/git-syncer/repository"
)

// NewConfigReadWriter mock config reader and writer
func NewConfigReadWriter() repository.ConfigReadWriter {
	c := make(map[string]string)
	return &mockConfigReadWriter{
		c,
	}
}

type mockConfigReadWriter struct {
	configs map[string]string
}

func (m *mockConfigReadWriter) GetConfig(keys ...string) (string, error) {
	for _, key := range keys {
		v, isExisted := m.configs[key]
		if isExisted {
			return v, nil
		}
	}

	return "", fmt.Errorf("config not found: keys=%v", keys)
}

func (m *mockConfigReadWriter) SetConfig(key, value string) error {
	m.configs[key] = value

	return nil
}
