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

func TestCatchup(t *testing.T) {
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
			err := syncer.Catchup(tc.contrib, tc.gitter)
			assert.Error(t, err)
			assert.Equal(t, tc.exitcode, errors.GetErrorCode(err))
		}
	})

	t.Run("failed to get repo head sha1", func(t *testing.T) {
		git := &gittertest.MockGitter{
			StubIn: gittertest.StubIn{
				GetHeadSHA1Return:    "",
				GetHeadSHA1ReturnErr: errors.Err(exitcode.RepoInvalidGitVersion, "mock invalid git version"),
			},
		}

		cb := &contribtest.MockContrib{}

		err := syncer.Catchup(cb, git)
		assert.Error(t, err)
		assert.Equal(t, exitcode.RepoInvalidGitVersion, errors.GetErrorCode(err))

	})

	t.Run("sync repo head sha1 to contrib", func(t *testing.T) {
		repoHeadSHA1 := "xxxxoooo"

		git := &gittertest.MockGitter{
			StubIn: gittertest.StubIn{
				GetHeadSHA1Return:    repoHeadSHA1,
				GetHeadSHA1ReturnErr: nil,
			},
		}

		t.Run("failed to sync", func(t *testing.T) {
			cb := &contribtest.MockContrib{
				StubIn: contribtest.StubIn{
					SyncReturnErr: errors.Err(exitcode.ContribLocked, "mock contrib locked"),
				},
			}

			err := syncer.Catchup(cb, git)
			assert.Error(t, err)
			assert.Equal(t, exitcode.ContribLocked, errors.GetErrorCode(err))

			assert.Equal(t, 1, cb.SyncCallTimes)
			assert.Equal(t, repoHeadSHA1, cb.SyncCallSHA1[0])
			assert.Nil(t, cb.SyncCallUploads[0], "catchup won't send uploads")
			assert.Nil(t, cb.SyncCallDeletes[0], "catchup won't send deletes")
		})

		t.Run("sync completed", func(t *testing.T) {
			cb := &contribtest.MockContrib{
				StubIn: contribtest.StubIn{
					SyncReturnErr: nil,
				},
			}

			err := syncer.Catchup(cb, git)
			assert.Nil(t, err)

			assert.Equal(t, 1, cb.SyncCallTimes)
			assert.Equal(t, repoHeadSHA1, cb.SyncCallSHA1[0])
			assert.Nil(t, cb.SyncCallUploads[0], "catchup won't send uploads")
			assert.Nil(t, cb.SyncCallDeletes[0], "catchup won't send deletes")
		})
	})
}
