package bucket

import (
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

// Bucket interface to access alioss bucket
type Bucket interface {
	GetObject(key string, options ...oss.Option) (io.ReadCloser, error)
	PutObject(key string, reader io.Reader, options ...oss.Option) error
	DeleteObject(key string, options ...oss.Option) error
	IsObjectExist(key string, options ...oss.Option) (bool, error)
	PutSymlink(symObjectKey string, targetObjectKey string, options ...oss.Option) error
}

// New create bucket
func New(opt *Options) (Bucket, error) {
	if opt == nil {
		return nil, errors.NewError(errors.WithMsg("bucket options requried"), errors.WithCode(exitcode.MissingArguments))
	}

	if optInvalid := opt.Valid(); optInvalid != nil {
		return nil, optInvalid
	}

	client, clientErr := oss.New(strings.TrimSpace(opt.Endpoint), opt.AccessKeyID, opt.AccessKeySecret)
	if clientErr != nil {
		return nil, clientErr
	}

	bkt, bucketErr := client.Bucket(strings.TrimSpace(opt.Bucket))
	if bucketErr != nil {
		return nil, bucketErr
	}

	return &bucket{
		opt,
		bkt,
	}, nil
}

type bucket struct {
	opt *Options
	*oss.Bucket
}
