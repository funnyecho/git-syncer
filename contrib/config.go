package contrib

import (
	"fmt"

	"github.com/funnyecho/git-syncer/repository"
)

// PrefixConfig config prefix
type PrefixConfig string

// MustConfig get config with prefix
func (p PrefixConfig) MustConfig(r repository.ConfigReader, key string) (string, error) {
	if p != "" {
		key = fmt.Sprintf("%s.%s", p, key)
	}

	return r.GetConfig(key)
}

// MayConfig get config with prefix, if error occur, return ""
func (p PrefixConfig) MayConfig(r repository.ConfigReader, key string) string {
	v, _ := p.MustConfig(r, key)
	return v
}
