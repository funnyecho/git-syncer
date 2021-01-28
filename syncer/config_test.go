package syncer_test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer"
	"github.com/funnyecho/git-syncer/syncer/gitter"
	"github.com/funnyecho/git-syncer/syncer/gitter/gittertest"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("failed with invalid params", func(t *testing.T) {
		git := &gittertest.MockGitter{}
		tcs := []struct {
			name string
			git  gitter.Gitter
			args []string
			code int
		}{
			{
				"failed without gitter",
				nil,
				[]string{"foo", "bar"},
				exitcode.InvalidParams,
			},
			{
				"failed without args",
				git,
				nil,
				exitcode.InvalidParams,
			},
			{
				"failed with empty args",
				git,
				[]string{},
				exitcode.Usage,
			},
			{
				"failed when length of args is greater than 2",
				git,
				[]string{"foo", "bar", "zoo"},
				exitcode.Usage,
			},
		}

		for _, tc := range tcs {
			err := syncer.Config(tc.git, tc.args)
			assert.Error(t, err)
			assert.Equal(t, tc.code, errors.GetErrorCode(err))
		}
	})

	t.Run("get config", func(t *testing.T) {

		t.Run("call gitter get config with key", func(t *testing.T) {
			git := &gittertest.MockGitter{}
			_ = syncer.Config(git, []string{"foo"})

			assert.Equal(t, 1, git.GetConfigCallTimes)
			assert.Equal(t, "foo", git.GetConfigCallKeys[0])
		})

		t.Run("failed when config not found or error occured", func(t *testing.T) {
			git := &gittertest.MockGitter{}

			t.Run("failed when gitter return empty string without error", func(t *testing.T) {
				git.GetConfigReturn = ""
				git.GetConfigReturnErr = nil

				err := syncer.Config(git, []string{"foo"})
				assert.Error(t, err)
				assert.Equal(t, exitcode.RepoConfigNotFound, errors.GetErrorCode(err))
			})

			t.Run("failed when gitter return with error", func(t *testing.T) {
				git.GetConfigReturn = ""
				git.GetConfigReturnErr = errors.Err(exitcode.RepoInvalidGitVersion, "mock invalid git version")

				err := syncer.Config(git, []string{"foo"})
				assert.Error(t, err)
				assert.Equal(t, exitcode.RepoInvalidGitVersion, errors.GetErrorCode(err), "return the error code from gitter")
			})
		})

		t.Run("println config", func(t *testing.T) {
			git := &gittertest.MockGitter{}
			git.GetConfigReturn = "bar"
			git.GetConfigReturnErr = nil

			pr, pw, _ := os.Pipe()

			go func() {
				defer func() {
					pw.Close()
				}()

				preStdout := os.Stdout

				os.Stdout = pw
				defer func() {
					os.Stdout = preStdout
				}()

				err := syncer.Config(git, []string{"foo"})
				assert.Nil(t, err)
			}()

			sb := new(strings.Builder)

			io.Copy(sb, pr)

			s := sb.String()

			assert.Equal(t, "bar\n", s)
		})
	})

	t.Run("set config", func(t *testing.T) {

		t.Run("call gitter set config with key and value", func(t *testing.T) {
			git := &gittertest.MockGitter{}

			_ = syncer.Config(git, []string{"foo", "bar"})

			assert.Equal(t, 1, git.SetConfigCallTimes)
			assert.Equal(t, "foo", git.SetConfigCallKeys[0])
			assert.Equal(t, "bar", git.SetConfigCallValues[0])
		})

		t.Run("failed with error occured", func(t *testing.T) {
			git := &gittertest.MockGitter{}
			git.SetConfigReturnErr = errors.Err(exitcode.RepoInvalidGitVersion, "mock invalid git version")

			err := syncer.Config(git, []string{"foo", "bar"})
			assert.Error(t, err)
			assert.Equal(t, exitcode.RepoInvalidGitVersion, errors.GetErrorCode(err), "return the error code from gitter")
		})

		t.Run("complete without error occured", func(t *testing.T) {
			git := &gittertest.MockGitter{}
			git.SetConfigReturnErr = nil

			err := syncer.Config(git, []string{"foo", "bar"})
			assert.Nil(t, err)
		})
	})
}
