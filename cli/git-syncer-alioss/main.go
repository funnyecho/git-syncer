package main

import (
	"github.com/funnyecho/git-syncer/adapter/contrib/alioss"
	"github.com/funnyecho/git-syncer/adapter/contrib/alioss/bucket"
	"github.com/funnyecho/git-syncer/adapter/contrib/alioss/options"
	"github.com/funnyecho/git-syncer/command"
	"github.com/funnyecho/git-syncer/syncer/contrib"
	"github.com/funnyecho/git-syncer/syncer/gitter"
)

func main() {
	contrib.WithFactory(contribFactory())
	command.Exec()
}

// NewFactory create alioss remote factory
func contribFactory() contrib.Factory {
	return func(params interface{ gitter.ConfigReader }) (contrib.Contrib, error) {
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
