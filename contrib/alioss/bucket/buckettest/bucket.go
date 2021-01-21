package buckettest

import (
	"fmt"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib/alioss/bucket"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

// New new mock bucket
func New(m Mocking) (bucket.Bucket, error) {
	return &Mock{
		&m,
	}, nil
}

// Mocking mocking to bucket interface
type Mocking struct {
	GetObject     func(key string, options ...oss.Option) (io.ReadCloser, error)
	PutObject     func(key string, reader io.Reader, options ...oss.Option) error
	DeleteObject  func(key string, options ...oss.Option) error
	IsObjectExist func(key string, options ...oss.Option) (bool, error)
	PutSymlink    func(symObjectKey string, targetObjectKey string, options ...oss.Option) error
}

// Mock bucket mocking
type Mock struct {
	mocking *Mocking
}

// GetObject GetObject
func (m *Mock) GetObject(key string, options ...oss.Option) (io.ReadCloser, error) {
	if m.mocking.GetObject == nil {
		return nil, fmt.Errorf("get object failed")
	}

	return m.mocking.GetObject(key, options...)
}

// PutObject PutObject
func (m *Mock) PutObject(key string, reader io.Reader, options ...oss.Option) error {
	if m.mocking.PutObject == nil {
		return errors.NewError(errors.WithCode(exitcode.ContribSyncFailed))
	}

	return m.mocking.PutObject(key, reader, options...)
}

// DeleteObject DeleteObject
func (m *Mock) DeleteObject(key string, options ...oss.Option) error {
	if m.mocking.DeleteObject == nil {
		return fmt.Errorf("delete object failed")
	}

	return m.mocking.DeleteObject(key, options...)
}

// IsObjectExist IsObjectExist
func (m *Mock) IsObjectExist(key string, options ...oss.Option) (bool, error) {
	if m.mocking.IsObjectExist == nil {
		return false, nil
	}

	return m.mocking.IsObjectExist(key, options...)
}

// PutSymlink PutSymlink
func (m *Mock) PutSymlink(symObjectKey string, targetObjectKey string, options ...oss.Option) error {
	if m.mocking.PutSymlink == nil {
		return fmt.Errorf("PutSymlink failed")
	}

	return m.mocking.PutSymlink(symObjectKey, targetObjectKey, options...)
}
