package options

import (
	"fmt"

	"github.com/funnyecho/git-syncer/syncer/gitter"
)

const (
	ossPrefixConfig = "alioss"
)

// New new Options
func New(r gitter.ConfigReader) *Options {
	return &Options{r}
}

// Options alioss Options
type Options struct {
	gitter.ConfigReader
}

// Endpoint endpoint options
func (o *Options) Endpoint() string {
	return o.mayConfig(o.prefixKey("endpoint"))
}

// AccessKeyID access key id options
func (o *Options) AccessKeyID() string {
	return o.mayConfig(o.prefixKey("access-key-id"))
}

// AccessKeySecret access key secret options
func (o *Options) AccessKeySecret() string {
	return o.mayConfig(o.prefixKey("access-key-secret"))
}

// Bucket bucket options
func (o *Options) Bucket() string {
	return o.mayConfig(o.prefixKey("bucket"))
}

// Base base options
func (o *Options) Base() string {
	return o.mayConfig(o.prefixKey("base"))
}

func (o *Options) mayConfig(key string) string {
	v, err := o.GetConfig(key)
	if err != nil {
		v = ""
	}

	return v
}

func (o *Options) prefixKey(key string) string {
	return fmt.Sprintf("%s.%s", ossPrefixConfig, key)
}
