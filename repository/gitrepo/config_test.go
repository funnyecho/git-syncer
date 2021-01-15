package gitrepo_test

import (
	"fmt"
	"testing"

	"github.com/funnyecho/git-syncer/repository/gitrepo"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter/gittertest"
	"github.com/stretchr/testify/assert"
)

func TestConfigReader(t *testing.T) {
	tcs := []struct {
		name        string
		remote      string
		key         string
		hit         func(key string) (string, error)
		expectValue string
		expectErred bool
	}{
		{
			name:   "hit value in remote segment with remote provided",
			remote: "test",
			key:    "foo",
			hit: func(key string) (string, error) {
				if key == "test.foo" {
					return "return test.foo", nil
				}

				assert.Fail(t, "shall only step into remote segment closure")
				return "", fmt.Errorf("config not found")
			},
			expectValue: "return test.foo",
			expectErred: false,
		},
		{
			name:   "hit value in basic segment with remote provided",
			remote: "test",
			key:    "foo",
			hit: func(key string) (string, error) {
				if key != "foo" && key != "test.foo" {
					assert.Fail(t, "shall only step into `foo` or `test.foo` statement")
				}

				if key == "test.foo" {
					return "", nil
				}

				if key == "foo" {
					return "return foo", nil
				}

				return "", fmt.Errorf("config not found")
			},
			expectValue: "return foo",
			expectErred: false,
		},
		{
			name: "hit value in basic segment without remote provided",
			key:  "foo",
			hit: func(key string) (string, error) {
				if key == "foo" {
					return "return foo", nil
				}

				assert.Fail(t, "shall only step into `foo` statement")
				return "", fmt.Errorf("config not found")
			},
			expectValue: "return foo",
			expectErred: false,
		},
		{
			name:   "config not found without remote provided",
			remote: "test",
			key:    "foo",
			hit: func(key string) (string, error) {
				if key != "foo" && key != "test.foo" {
					assert.Fail(t, "shall only step into `foo` or `test.foo` statement")
				}

				return "", fmt.Errorf("something failed")
			},
			expectValue: "",
			expectErred: true,
		},
		{
			name: "config not found without remote provided",
			key:  "foo",
			hit: func(key string) (string, error) {
				if key != "foo" {
					assert.Fail(t, "shall only step into `foo` statement")
				}

				return "", fmt.Errorf("something failed")
			},
			expectValue: "",
			expectErred: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			repo, _ := gitrepo.New(
				gitrepo.WithWorkingRemote(tc.remote),
				gitrepo.WithGitter(&configTestGitter{
					gittertest.New(),
					func(key string, options gitter.ConfigGetOptions) (string, error) {
						assert.Equal(t, gitrepo.ProjectConfigName, options.File)

						return tc.hit(key)
					},
				}),
			)

			v, err := repo.GetConfig(tc.key)
			if tc.expectErred {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectValue, v)
		})
	}
}

func TestConfigWriter(t *testing.T) {

}

type configTestGitter struct {
	gitter.Gitter
	configGet func(key string, options gitter.ConfigGetOptions) (string, error)
}

func (g *configTestGitter) ConfigGet(key string, options gitter.ConfigGetOptions) (string, error) {
	return g.configGet(key, options)
}
