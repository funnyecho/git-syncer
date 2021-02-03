package main

import (
	"github.com/funnyecho/git-syncer/adapter/remote/alioss"
	"github.com/funnyecho/git-syncer/adapter/remote/alioss/bucket"
	"github.com/funnyecho/git-syncer/adapter/remote/alioss/options"
	"github.com/funnyecho/git-syncer/command"
	"github.com/funnyecho/git-syncer/syncer/gitter"
	"github.com/funnyecho/git-syncer/syncer/remote"
)

func main() {
	remote.WithFactory(remoteFactory())
	command.Exec()
}

// NewFactory create alioss remote factory
func remoteFactory() remote.Factory {
	return func(params interface{ gitter.ConfigReader }) (remote.Remote, error) {
		ossOptions := options.New(params)

		bucket, bucketErr := bucket.New(&bucket.Options{
			Endpoint:        ossOptions.Endpoint(),
			Bucket:          ossOptions.Bucket(),
			AccessKeyID:     ossOptions.AccessKeyID(),
			AccessKeySecret: ossOptions.AccessKeySecret(),
		})
		if bucketErr != nil {
			return nil, bucketErr
		}

		return alioss.New(ossOptions, bucket)
	}
}
