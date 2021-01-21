package repotest

import "github.com/funnyecho/git-syncer/repository"

// NewHeadReader mock head reader
func NewHeadReader() *MockHeadReader {
	return &MockHeadReader{}
}

// NewHeadWriter mock head writer
func NewHeadWriter() *MockHeadWriter {
	return &MockHeadWriter{}
}

// NewHeadReadWriter mock head reader and writer
func NewHeadReadWriter() *MockHeadReadWriter {
	return &MockHeadReadWriter{
		NewHeadReader(),
		NewHeadWriter(),
	}
}

// MockHeadReader mock head reader
type MockHeadReader struct {
	Head string
}

// MockHeadWriter mock head writer
type MockHeadWriter struct {
	Head string
}

// MockHeadReadWriter mock head read and writer
type MockHeadReadWriter struct {
	repository.HeadReader
	repository.HeadWriter
}

// GetHead get head
func (m *MockHeadReader) GetHead() (string, error) {
	return m.Head, nil
}

// GetHeadSHA1 get head sha1
func (m *MockHeadReader) GetHeadSHA1() (string, error) {
	return "", nil
}

// PushHead push head
func (m *MockHeadWriter) PushHead(head string) (string, error) {
	defer func() {
		m.Head = head
	}()
	return m.Head, nil
}

// PopHead pop head
func (m *MockHeadWriter) PopHead(head string) error {
	m.Head = head
	return nil
}
