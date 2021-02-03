package buckettest

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// MockBucket mock bucket
type MockBucket struct {
	StubIn
	StubOut
}

// GetObject mock GetObject
func (m *MockBucket) GetObject(key string, options ...oss.Option) (io.ReadCloser, error) {
	m.GetObjectCallTimes++
	m.GetObjectCallKey = append(m.GetObjectCallKey, key)
	m.GetObjectCallOptions = append(m.GetObjectCallOptions, options)

	return m.GetObjectReturn, m.GetObjectReturnErr
}

// PutObject mock PutObject
func (m *MockBucket) PutObject(key string, reader io.Reader, options ...oss.Option) error {
	m.PutObjectCallTimes++
	m.PutObjectCallKey = append(m.PutObjectCallKey, key)
	m.PutObjectCallReader = append(m.PutObjectCallReader, reader)
	m.PutObjectCallOptions = append(m.PutObjectCallOptions, options)

	return m.PutObjectReturnErr
}

// DeleteObject mock DeleteObject
func (m *MockBucket) DeleteObject(key string, options ...oss.Option) error {
	m.DeleteObjectCallTimes++
	m.DeleteObjectCallKey = append(m.DeleteObjectCallKey, key)
	m.DeleteObjectCallOptions = append(m.DeleteObjectCallOptions, options)

	return m.DeleteObjectReturnErr
}

// IsObjectExist mock IsObjectExist
func (m *MockBucket) IsObjectExist(key string, options ...oss.Option) (bool, error) {
	m.IsObjectExistCallTimes++
	m.IsObjectExistCallKey = append(m.IsObjectExistCallKey, key)
	m.IsObjectExistCallOptions = append(m.IsObjectExistCallOptions, options)

	return m.IsObjectExistReturn, m.IsObjectExistReturnErr
}

// PutSymlink mock PutSymlink
func (m *MockBucket) PutSymlink(symObjectKey string, targetObjectKey string, options ...oss.Option) error {
	m.PutSymlinkCallTimes++
	m.PutSymlinkCallSymKey = append(m.PutSymlinkCallSymKey, symObjectKey)
	m.PutSymlinkCallTargetKey = append(m.PutSymlinkCallTargetKey, targetObjectKey)
	m.PutSymlinkCallOptions = append(m.PutSymlinkCallOptions, options)

	return m.PutSymlinkReturnErr
}
