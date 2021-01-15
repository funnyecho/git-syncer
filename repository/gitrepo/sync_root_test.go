package gitrepo_test

import (
	"testing"

	"github.com/funnyecho/git-syncer/repository/gitrepo"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter/gittertest"
	"github.com/stretchr/testify/assert"
)

func TestGetSyncRoot(t *testing.T) {
	tcs := []struct {
		name      string
		configGet func(key string, options gitter.ConfigGetOptions) (string, error)
		expectVal string
	}{
		{
			"return non-empty config from gitter",
			func(key string, options gitter.ConfigGetOptions) (string, error) {
				return "./foo/bar", nil
			},
			"./foo/bar",
		},
		{
			"return default sync root when config is empty",
			func(key string, options gitter.ConfigGetOptions) (string, error) {
				return "", nil
			},
			gitrepo.DefaultSyncRoot,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			repo, _ := gitrepo.New(
				gitrepo.WithGitter(&configTestGitter{
					gittertest.New(),
					tc.configGet,
				}),
			)

			v := gitrepo.GetSyncRoot(repo)
			assert.Equal(t, tc.expectVal, v)
		})
	}
}
