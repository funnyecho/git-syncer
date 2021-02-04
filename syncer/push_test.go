package syncer_test

import (
	"testing"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer"
	"github.com/funnyecho/git-syncer/syncer/contrib"
	"github.com/funnyecho/git-syncer/syncer/contrib/contribtest"
	"github.com/funnyecho/git-syncer/syncer/gitter"
	"github.com/funnyecho/git-syncer/syncer/gitter/gittertest"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	t.Run("invalid params", func(t *testing.T) {
		tcs := []struct {
			contrib  contrib.Contrib
			gitter   gitter.Gitter
			exitcode int
		}{
			{
				nil,
				nil,
				exitcode.InvalidParams,
			},
			{
				nil,
				nil,
				exitcode.InvalidParams,
			},
			{
				&contribtest.MockContrib{},
				nil,
				exitcode.InvalidParams,
			},
		}

		for _, tc := range tcs {
			err := syncer.Push(tc.contrib, tc.gitter)
			assert.Error(t, err)
			assert.Equal(t, tc.exitcode, errors.GetErrorCode(err))
		}
	})

	tcs := []struct {
		name          string
		contribStubIn contribtest.StubIn
		gitterStubIn  gittertest.StubIn
		exitcode      int
	}{
		{
			name: "failed to get contrib head sha1",
			contribStubIn: contribtest.StubIn{
				GetHeadSHA1Return:    "",
				GetHeadSHA1ReturnErr: errors.Err(exitcode.ContribInvalidLog, "mock contrib log invalid"),
			},
			exitcode: exitcode.ContribInvalidLog,
		},
		{
			name: "failed when contrib head sha1 is empty",
			contribStubIn: contribtest.StubIn{
				GetHeadSHA1Return:    "",
				GetHeadSHA1ReturnErr: nil,
			},
			exitcode: exitcode.ContribHeadNotFound,
		},
		{
			name: "failed to get changed files from repo",
			contribStubIn: contribtest.StubIn{
				GetHeadSHA1Return:    "xxooxxoo",
				GetHeadSHA1ReturnErr: nil,
			},
			gitterStubIn: gittertest.StubIn{
				ListChangedFilesReturnErr: errors.Err(exitcode.RepoDiffBaseNotFound, "mock repo diss base not found"),
			},
			exitcode: exitcode.RepoDiffBaseNotFound,
		},
		{
			name: "failed to sync changed files to contrib",
			contribStubIn: contribtest.StubIn{
				GetHeadSHA1Return:    "xxooxxoo",
				GetHeadSHA1ReturnErr: nil,

				SyncReturnErr: errors.Err(exitcode.ContribLocked, "mock contrib locked"),
			},
			exitcode: exitcode.ContribLocked,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			cb := &contribtest.MockContrib{
				StubIn: tc.contribStubIn,
			}

			git := &gittertest.MockGitter{
				StubIn: tc.gitterStubIn,
			}

			err := syncer.Push(cb, git)
			assert.Error(t, err)
			assert.Equal(t, tc.exitcode, errors.GetErrorCode(err))
		})
	}

	cb := &contribtest.MockContrib{
		StubIn: contribtest.StubIn{
			GetHeadSHA1Return: "ooxxooxx",
			SyncReturnUploaded: func(s []string) []string {
				return s
			},
			SyncReturnDeleted: func(s []string) []string {
				return s
			},
		},
	}

	git := &gittertest.MockGitter{
		StubIn: gittertest.StubIn{
			GetConfigReturnFn: func(s string) (string, error) {
				if s == "sync-root" {
					return "milan city", nil
				}

				return "", errors.Err(exitcode.RepoConfigNotFound, "config not found")
			},
			GetHeadSHA1Return:        "xxooxxoo",
			ListChangedFilesReturnAM: []string{"foo", "bar"},
			ListChangedFilesReturnD:  []string{"zoo"},
		},
	}

	err := syncer.Push(cb, git)
	assert.Nil(t, err)

	assert.Equal(t, 1, cb.GetHeadSHA1CallTimes)
	assert.Equal(t, 1, cb.SyncCallTimes)
	assert.Equal(t, "xxooxxoo", cb.SyncCallSHA1[0])
	assert.Equal(t, []string{"foo", "bar"}, cb.SyncCallUploads[0])
	assert.Equal(t, []string{"zoo"}, cb.SyncCallDeletes[0])

	assert.Equal(t, 1, git.GetHeadSHA1CallTimes)
	assert.Equal(t, 1, git.ListChangedFilesCallTimes)
	assert.Equal(t, "milan city", git.ListChangedFilesCallPaths[0])
	assert.Equal(t, "ooxxooxx", git.ListChangedFilesCallBaseSHA1[0])

}
