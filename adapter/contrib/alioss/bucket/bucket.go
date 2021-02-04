package bucket

import (
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

// New create bucket
func New(opt *Options) (*Bucket, error) {
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

	return &Bucket{
		opt,
		bkt,
	}, nil
}

type Bucket struct {
	opt *Options
	*oss.Bucket
}
