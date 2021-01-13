package gitrepo_test

import (
	"fmt"
	"testing"

	"github.com/funnyecho/git-syncer/repository/gitrepo"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter/gitter_test"
	"github.com/stretchr/testify/assert"
)

func TestIsDirtyRepository(t *testing.T) {
	tcs := []struct {
		name                  string
		getUnoPorcelainStatus func() (status []string, err error)
		expectErred           bool
		expectVal             bool
	}{
		{
			"failed to get repo uno porcelain status",
			func() (status []string, err error) {
				return nil, fmt.Errorf("failed to get repo uno porcelain status")
			},
			true,
			false,
		},
		{
			"dirty when uno porcelain status is not empty",
			func() (status []string, err error) {
				return []string{"foo", "bar"}, nil
			},
			false,
			false,
		},
		{
			"clean when uno porcelain status is empty",
			func() (status []string, err error) {
				return []string{}, nil
			},
			false,
			true,
		},
		{
			"clean when uno porcelain status is nil",
			func() (status []string, err error) {
				return nil, nil
			},
			false,
			true,
		},
	}

	for _, tc := range tcs {
		gitter := &filesTestGitter{
			Gitter:                gitter_test.New(),
			getUnoPorcelainStatus: tc.getUnoPorcelainStatus,
		}

		v, err := gitrepo.IsDirtyRepository(gitter)
		if tc.expectErred {
			assert.Error(t, err)
			return
		}

		assert.Equal(t, tc.expectVal, v)
	}
}
