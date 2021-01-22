package command_test

import (
	stdErr "errors"
	"testing"

	"github.com/funnyecho/git-syncer/command"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/contrib/contribtest"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	tcs := []struct {
		name            string
		contribHeadSha1 contribtest.ContribHeadSha1Fetcher
		contribSyncFile contribtest.ContribSyncHandler
		repoListFiles   contribtest.RepoListAllFiles
		err             error
	}{
		{
			name: "failed to get deployed head sha1",
			contribHeadSha1: func() (string, error) {
				return "", errors.NewError(
					errors.WithCode(exitcode.ContribForbidden),
				)
			},
			err: errors.NewError(
				errors.WithCode(exitcode.ContribForbidden),
			),
		},
		{
			name: "failed to setup when deployed head sha1 existed",
			contribHeadSha1: func() (string, error) {
				return "contrib-sha1", nil
			},
			err: errors.NewError(
				errors.WithCode(exitcode.Usage),
			),
		},
		{
			name: "failed to get repo files",
			repoListFiles: func() (sha1 string, uploads []string, err error) {
				return "", nil, errors.NewError(
					errors.WithCode(exitcode.RepoListFilesFailed),
				)
			},
			err: errors.NewError(
				errors.WithCode(exitcode.RepoListFilesFailed),
			),
		},
		{
			name: "setup failed when upload failed",
			repoListFiles: func() (sha1 string, uploads []string, err error) {
				return "Foobar", []string{"foo", "bar", "zoo"}, nil
			},
			contribSyncFile: func(req *contrib.SyncReq) (contrib.SyncRes, error) {
				return contrib.SyncRes{}, errors.NewError(errors.WithCode(exitcode.ContribSyncFailed))
			},
			err: errors.NewError(errors.WithCode(exitcode.ContribSyncFailed)),
		},
	}

	for _, tc := range tcs {
		cb := &contribtest.MockContrib{
			HeadSHA1:    tc.contribHeadSha1,
			SyncHandler: tc.contribSyncFile,
		}

		rp := &mockSetupRepo{
			listFiles: tc.repoListFiles,
		}

		t.Run(tc.name, func(t *testing.T) {
			e := command.ExecSetup(cb, rp)
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

		cb := &contribtest.MockContrib{
			SyncHandler: func(req *contrib.SyncReq) (contrib.SyncRes, error) {
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

		e := command.ExecSetup(cb, rp)
		assert.Nil(t, e)
	})
}

type mockSetupRepo struct {
	listFiles contribtest.RepoListAllFiles
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
