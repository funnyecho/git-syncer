package contrib

import (
	"fmt"

	"github.com/funnyecho/git-syncer/repository"
)

// Configurable prefix wrapper for config
type Configurable struct {
	prefix string
	reader repository.ConfigReader
}

// NewConfigurable create configurable instance
func NewConfigurable(prefix string, reader repository.ConfigReader) *Configurable {
	return &Configurable{
		prefix,
		reader,
	}
}

// GetConfig get config with contrib prefix from repo
func (c *Configurable) GetConfig(key string) (string, error) {
	if c.prefix != "" {
		key = fmt.Sprintf("%s.%s", c.prefix, key)
	}

	return c.reader.GetConfig(key)
}

// GetRawConfig get config without prefix from repo
func (c *Configurable) GetRawConfig(key string) (string, error) {
	return c.reader.GetConfig(key)
}
