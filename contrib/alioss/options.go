package alioss

import (
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/repository"
)

const (
	ossPrefixConfig  = contrib.PrefixConfig("alioss")
	noopPrefixConfig = contrib.PrefixConfig("")
)

// NewOptions new Options
func NewOptions(r repository.ConfigReader) *Options {
	return &Options{r}
}

// Options alioss Options
type Options struct {
	repository.ConfigReader
}

// Endpoint endpoint options
func (o *Options) Endpoint() string {
	return o.ossMayConfig("endpoint")
}

// AccessKeyID access key id options
func (o *Options) AccessKeyID() string {
	return o.ossMayConfig("access_key_id")
}

// AccessKeySecret access key secret options
func (o *Options) AccessKeySecret() string {
	return o.ossMayConfig("access_key_secret")
}

// Bucket bucket options
func (o *Options) Bucket() string {
	return o.ossMayConfig("bucket")
}

// Base base options
func (o *Options) Base() string {
	return o.ossMayConfig("base")
}

// UserName user.name
func (o *Options) UserName() string {
	return o.noopMayConfig("user.name")
}

// UserEmail user.email
func (o *Options) UserEmail() string {
	return o.noopMayConfig("user.email")
}

func (o *Options) ossMayConfig(key string) string {
	return ossPrefixConfig.MayConfig(o.ConfigReader, key)
}

func (o *Options) noopMayConfig(key string) string {
	return noopPrefixConfig.MayConfig(o.ConfigReader, key)
}
