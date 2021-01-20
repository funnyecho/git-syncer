package alioss

import "github.com/funnyecho/git-syncer/contrib"

// Options options to setup alioss contrib
type Options struct {
	Endpoint        string
	Bucket          string
	AccessKeyID     string
	AccessKeySecret string
	Base            string
}

// NewOptions init options from config
func NewOptions(c *contrib.Configurable) (*Options, error) {
	endpoint, endpointErr := c.GetConfig("endpoint")
	if endpointErr != nil {
		return nil, endpointErr
	}

	bucket, bucketErr := c.GetConfig("bucket")
	if bucketErr != nil {
		return nil, bucketErr
	}

	akID, akIDErr := c.GetConfig("access_key_id")
	if akIDErr != nil {
		return nil, akIDErr
	}

	akSecret, akSecretErr := c.GetConfig("access_key_secret")
	if akSecretErr != nil {
		return nil, akSecretErr
	}

	base, baseErr := c.GetConfig("base")
	if baseErr != nil {
		return nil, baseErr
	}

	return &Options{
		endpoint,
		bucket,
		akID,
		akSecret,
		base,
	}, nil
}
