package bucket

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

// Options options to create oss client and access bucket
type Options struct {
	Endpoint        string
	Bucket          string
	AccessKeyID     string
	AccessKeySecret string
}

// Valid check options valid
func (o *Options) Valid() error {
	if o.Endpoint == "" || o.Bucket == "" || o.AccessKeyID == "" || o.AccessKeySecret == "" {
		return errors.Err(
			exitcode.MissingArguments,
			"invalid alioss bucket options, endpoint, bucket, accessKeyID and accessKeySecret are required",
		)
	}

	return nil
}
