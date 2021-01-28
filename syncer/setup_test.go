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

func TestSetup(t *testing.T) {
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
			err := syncer.Setup(tc.remote, tc.gitter)
			assert.Error(t, err)
			assert.Equal(t, tc.exitcode, errors.GetErrorCode(err))
		}
	})

	tcs := []struct {
		name         string
		remoteStubIn remotetest.StubIn
		gitterStubIn gittertest.StubIn
		exitcode     int
	}{
		{
			name: "failed to get remote head sha1",
			remoteStubIn: remotetest.StubIn{
				GetHeadSHA1Return:    "",
				GetHeadSHA1ReturnErr: errors.Err(exitcode.RemoteInvalidLog, "mock remote log invalid"),
			},
			exitcode: exitcode.RemoteInvalidLog,
		},
		{
			name: "failed when remote head sha1 is not empty",
			remoteStubIn: remotetest.StubIn{
				GetHeadSHA1Return:    "ooxxooxx",
				GetHeadSHA1ReturnErr: nil,
			},
			exitcode: exitcode.Usage,
		},
		{
			name: "failed to get all tracked files from repo",
			gitterStubIn: gittertest.StubIn{
				ListTrackedFilesReturnErr: errors.Err(exitcode.RepoInvalidGitVersion, "mock invalid git version"),
			},
			exitcode: exitcode.RepoInvalidGitVersion,
		},
		{
			name: "failed to sync changed files to remote",
			remoteStubIn: remotetest.StubIn{
				SyncReturnErr: errors.Err(exitcode.RemoteLocked, "mock remote locked"),
			},
			exitcode: exitcode.RemoteLocked,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			rm := &remotetest.MockRemote{
				StubIn: tc.remoteStubIn,
			}

			git := &gittertest.MockGitter{
				StubIn: tc.gitterStubIn,
			}

			err := syncer.Setup(rm, git)
			assert.Error(t, err)
			assert.Equal(t, tc.exitcode, errors.GetErrorCode(err))
		})
	}

	rm := &remotetest.MockRemote{
		StubIn: remotetest.StubIn{
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
			GetHeadSHA1Return:      "xxooxxoo",
			ListTrackedFilesReturn: []string{"foo", "bar"},
		},
	}

	err := syncer.Setup(rm, git)
	assert.Nil(t, err)

	assert.Equal(t, 1, rm.GetHeadSHA1CallTimes)
	assert.Equal(t, 1, rm.SyncCallTimes)
	assert.Equal(t, "xxooxxoo", rm.SyncCallSHA1[0])
	assert.Equal(t, []string{"foo", "bar"}, rm.SyncCallUploads[0])
	assert.Nil(t, rm.SyncCallDeletes[0])

	assert.Equal(t, 1, git.GetHeadSHA1CallTimes)
	assert.Equal(t, 1, git.ListTrackedFilesCallTimes)
	assert.Equal(t, "milan city", git.ListTrackedFilesCallPaths[0])
}
