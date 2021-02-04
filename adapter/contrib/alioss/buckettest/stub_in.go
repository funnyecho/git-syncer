package buckettest

import "io"

// StubIn mock calling source
type StubIn struct {
	GetObjectReturn    io.ReadCloser
	GetObjectReturnErr error

	PutObjectReturnErr error

	DeleteObjectReturnErr error

	IsObjectExistReturn    bool
	IsObjectExistReturnErr error

	PutSymlinkReturnErr error
}
