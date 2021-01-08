package contrib_test

import (
	stdErr "errors"
	"testing"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	tcs := []struct {
		name            string
		contribHeadSha1 ContribHeadSha1Fetcher
		contribSyncFile ContribSyncHandler
		repoListFiles   RepoListAllFiles
		err             error
	}{
		{
			name: "failed to get deployed head sha1",
			contribHeadSha1: func() (string, error) {
				return "", errors.NewError(
					errors.WithStatusCode(exitcode.RemoteForbidden),
				)
			},
			err: errors.NewError(
				errors.WithStatusCode(exitcode.RemoteForbidden),
			),
		},
		{
			name: "failed to setup when deployed head sha1 existed",
			contribHeadSha1: func() (string, error) {
				return "contrib-sha1", nil
			},
			err: errors.NewError(
				errors.WithStatusCode(exitcode.Usage),
			),
		},
		{
			name: "failed to get repo files",
			repoListFiles: func() (sha1 string, uploads []string, err error) {
				return "", nil, errors.NewError(
					errors.WithStatusCode(exitcode.Git),
				)
			},
			err: errors.NewError(
				errors.WithStatusCode(exitcode.Git),
			),
		},
		{
			name: "setup failed when upload failed",
			repoListFiles: func() (sha1 string, uploads []string, err error) {
				return "Foobar", []string{"foo", "bar", "zoo"}, nil
			},
			contribSyncFile: func(req *contrib.SyncReq) (contrib.SyncRes, error) {
				return contrib.SyncRes{}, errors.NewError(errors.WithStatusCode(exitcode.Upload))
			},
			err: errors.NewError(errors.WithStatusCode(exitcode.Upload)),
		},
	}

	for _, tc := range tcs {
		cb := &mockSetupContrib{
			headSHA1:    tc.contribHeadSha1,
			syncHandler: tc.contribSyncFile,
		}
		rp := &mockSetupRepo{
			listFiles: tc.repoListFiles,
		}

		t.Run(tc.name, func(t *testing.T) {
			e := contrib.Setup(cb, rp)
			if tc.err == nil {
				assert.Nil(t, e)
			} else {
				assert.True(t, stdErr.Is(tc.err, e))
			}
		})
	}

	t.Run("repo files was uploaded, and repo sha1 was deployed", func(t *testing.T) {
		repoSha1 := "foobar"
		toBeUploads := []string{"foo", "bar", "zoo"}

		cb := &mockSetupContrib{
			syncHandler: func(req *contrib.SyncReq) (contrib.SyncRes, error) {
				assert.Equal(t, repoSha1, req.SHA1)
				assert.Equal(t, toBeUploads, req.Uploads)
				assert.Nil(t, req.Deletes)

				return contrib.SyncRes{}, nil
			},
		}
		rp := &mockSetupRepo{
			listFiles: func() (sha1 string, uploads []string, err error) {
				return repoSha1, toBeUploads, nil
			},
		}

		e := contrib.Setup(cb, rp)
		assert.Nil(t, e)
	})
}

type mockSetupRepo struct {
	listFiles RepoListAllFiles
}

func (r *mockSetupRepo) ListAllFiles() (sha1 string, uploads []string, err error) {
	if r.listFiles == nil {
		return "", nil, nil
	}

	return r.listFiles()
}

func (r *mockSetupRepo) ListChangedFiles(baseSha1 string) (sha1 string, uploads []string, deletes []string, err error) {
	panic("please implement me")
}

type mockSetupContrib struct {
	headSHA1    ContribHeadSha1Fetcher
	syncHandler ContribSyncHandler
}

func (c *mockSetupContrib) GetHeadSHA1() (string, error) {
	if c.headSHA1 == nil {
		return "", nil
	}

	return c.headSHA1()
}

func (c *mockSetupContrib) Sync(req *contrib.SyncReq) (contrib.SyncRes, error) {
	if c.syncHandler == nil {
		return contrib.SyncRes{}, nil
	}

	return c.syncHandler(req)
}
