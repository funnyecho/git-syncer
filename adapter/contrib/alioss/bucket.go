package alioss

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Bucket bucket interface to interact with alioss
type Bucket interface {
	GetObject(key string, options ...oss.Option) (io.ReadCloser, error)
	PutObject(key string, reader io.Reader, options ...oss.Option) error
	DeleteObject(key string, options ...oss.Option) error
	IsObjectExist(key string, options ...oss.Option) (bool, error)
	PutSymlink(symObjectKey string, targetObjectKey string, options ...oss.Option) error
}
