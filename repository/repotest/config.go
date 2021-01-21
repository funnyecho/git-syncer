package repotest

import (
	"fmt"
)

// NewConfigReadWriter mock config reader and writer
func NewConfigReadWriter() *MockConfigReadWriter {
	c := make(map[string]string)
	return &MockConfigReadWriter{
		c,
	}
}

// MockConfigReadWriter mock config reader and writer
type MockConfigReadWriter struct {
	Configs map[string]string
}

// GetConfig get config
func (m *MockConfigReadWriter) GetConfig(key string) (string, error) {
	v, isExisted := m.Configs[key]
	if !isExisted {
		return "", fmt.Errorf("config not found: key=%s", key)
	}

	return v, nil
}

// SetConfig update config
func (m *MockConfigReadWriter) SetConfig(key, value string) error {
	m.Configs[key] = value

	return nil
}
