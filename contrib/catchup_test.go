package contrib_test

import (
	"fmt"
	"testing"

	"github.com/funnyecho/git-syncer/contrib"
	"github.com/stretchr/testify/assert"
)

func TestCatchup(t *testing.T) {
	tcs := []struct {
		name          string
		repoHeadSHA1  RepoGetHeadSHA1
		contribSyncer ContribSyncHandler
		expectErred   bool
	}{
		{
			"failed to get repo head sha1",
			func() (string, error) {
				return "", fmt.Errorf("failed to get repo head sha1")
			},
			func(sr *contrib.SyncReq) (contrib.SyncRes, error) {
				assert.Fail(t, "shall not call contrib syncer")
				return contrib.SyncRes{}, nil
			},
			true,
		},
		{
			"repo head sha1 can't be empty",
			func() (string, error) {
				return "", nil
			},
			func(sr *contrib.SyncReq) (contrib.SyncRes, error) {
				assert.Fail(t, "shall not call contrib syncer")
				return contrib.SyncRes{}, nil
			},
			true,
		},
		{
			"failed to deployed sha1",
			func() (string, error) {
				return "abcd1234", nil
			},
			func(sr *contrib.SyncReq) (contrib.SyncRes, error) {
				assert.Nil(t, sr.Uploads)
				assert.Nil(t, sr.Deletes)
				return contrib.SyncRes{}, fmt.Errorf("failed to deployed sha1")
			},
			true,
		},
		{
			"sha1 deployed succefully",
			func() (string, error) {
				return "abcd1234", nil
			},
			func(sr *contrib.SyncReq) (contrib.SyncRes, error) {
				assert.Nil(t, sr.Uploads)
				assert.Nil(t, sr.Deletes)

				return contrib.SyncRes{
					"abcd1234",
					nil,
					nil,
				}, nil
			},
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := contrib.Catchup(
				&mockContrib{
					nil,
					tc.contribSyncer,
				},
				&mockCatchupRepo{
					tc.repoHeadSHA1,
				},
			)
			if tc.expectErred {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

type mockCatchupRepo struct {
	getHeadSHA1 RepoGetHeadSHA1
}

func (m *mockCatchupRepo) GetHead() (string, error) {
	return "", nil
}

func (m *mockCatchupRepo) GetHeadSHA1() (string, error) {
	if m.getHeadSHA1 != nil {
		return m.getHeadSHA1()
	}

	return "", nil
}
