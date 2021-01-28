package syncer_test

import (
	"testing"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer"
	"github.com/funnyecho/git-syncer/syncer/gitter"
	"github.com/funnyecho/git-syncer/syncer/gitter/gittertest"
	"github.com/funnyecho/git-syncer/syncer/remote"
	"github.com/funnyecho/git-syncer/syncer/remote/remotetest"
	"github.com/stretchr/testify/assert"
)

func TestCatchup(t *testing.T) {
	t.Run("invalid params", func(t *testing.T) {
		tcs := []struct {
			remote   remote.Remote
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
				&remotetest.MockRemote{},
				nil,
				exitcode.InvalidParams,
			},
		}

		for _, tc := range tcs {
			err := syncer.Catchup(tc.remote, tc.gitter)
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

		rm := &remotetest.MockRemote{}

		err := syncer.Catchup(rm, git)
		assert.Error(t, err)
		assert.Equal(t, exitcode.RepoInvalidGitVersion, errors.GetErrorCode(err))

	})

	t.Run("sync repo head sha1 to remote", func(t *testing.T) {
		repoHeadSHA1 := "xxxxoooo"

		git := &gittertest.MockGitter{
			StubIn: gittertest.StubIn{
				GetHeadSHA1Return:    repoHeadSHA1,
				GetHeadSHA1ReturnErr: nil,
			},
		}

		t.Run("failed to sync", func(t *testing.T) {
			rm := &remotetest.MockRemote{
				StubIn: remotetest.StubIn{
					SyncReturnErr: errors.Err(exitcode.RemoteLocked, "mock remote locked"),
				},
			}

			err := syncer.Catchup(rm, git)
			assert.Error(t, err)
			assert.Equal(t, exitcode.RemoteLocked, errors.GetErrorCode(err))

			assert.Equal(t, 1, rm.SyncCallTimes)
			assert.Equal(t, repoHeadSHA1, rm.SyncCallSHA1[0])
			assert.Nil(t, rm.SyncCallUploads[0], "catchup won't send uploads")
			assert.Nil(t, rm.SyncCallDeletes[0], "catchup won't send deletes")
		})

		t.Run("sync completed", func(t *testing.T) {
			rm := &remotetest.MockRemote{
				StubIn: remotetest.StubIn{
					SyncReturnErr: nil,
				},
			}

			err := syncer.Catchup(rm, git)
			assert.Nil(t, err)

			assert.Equal(t, 1, rm.SyncCallTimes)
			assert.Equal(t, repoHeadSHA1, rm.SyncCallSHA1[0])
			assert.Nil(t, rm.SyncCallUploads[0], "catchup won't send uploads")
			assert.Nil(t, rm.SyncCallDeletes[0], "catchup won't send deletes")
		})
	})
}
