package repository_test

import "github.com/funnyecho/git-syncer/repository"

// NewHeadReader mock head reader
func NewHeadReader() repository.HeadReader {
	return &mockHeadReader{}
}

// NewHeadWriter mock head writer
func NewHeadWriter() repository.HeadWriter {
	return &mockHeadWriter{}
}

// NewHeadReadWriter mock head reader and writer
func NewHeadReadWriter() repository.HeadReadWriter {
	return &mockHeadReadWriter{
		NewHeadReader(),
		NewHeadWriter(),
	}
}

type mockHeadReader struct {
}

type mockHeadWriter struct {
	head string
}

type mockHeadReadWriter struct {
	repository.HeadReader
	repository.HeadWriter
}

func (m *mockHeadReader) GetHead() (string, error) {
	return "", nil
}

func (m *mockHeadReader) GetHeadSHA1() (string, error) {
	return "", nil
}

func (m *mockHeadWriter) PushHead(head string) (string, error) {
	defer func() {
		m.head = head
	}()
	return m.head, nil
}

func (m *mockHeadWriter) PopHead(head string) error {
	m.head = head
	return nil
}
