package setup_test

import (
	"github.com/funnyecho/git-syncer/command/setup"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetupFailed(t *testing.T) {
	t.Skip()
	cmd, _ := setup.Factory()

	t.Run("Invalid git version", func(t *testing.T) {
		t.Run("git version < 1.7", func(t *testing.T) {
			code := cmd.Run(nil)
			assert.Equal(t, exitcode.Git, code)
		})
	})

	t.Run("Not inside a git repository", func(t *testing.T) {
		code := cmd.Run(nil)

		assert.Equal(t, exitcode.Git, code)
	})

	t.Run("Git repository is dirty", func(t *testing.T) {
		code := cmd.Run(nil)

		assert.Equal(t, exitcode.Git, code)
	})

	t.Run("Branch not existed", func(t *testing.T) {
		assert.Fail(t, "todo")
	})

	t.Run("Remote syncer contrib setup failed", func(t *testing.T) {
		assert.Fail(t, "todo")
	})

	t.Run("Remote ref existed", func(t *testing.T) {
		assert.Fail(t, "todo")
	})

	t.Run("Local syncer was locked", func(t *testing.T) {
		assert.Fail(t, "todo")
	})

	t.Run("Remote syncer was locked", func(t *testing.T) {
		assert.Fail(t, "todo")
	})
}

func TestSetupCompleted(t *testing.T) {
	t.Skip()

	t.Run("Local syncer haven't been locked", func(t *testing.T) {
		assert.Fail(t, "todo")
	})

	t.Run("Remove syncer haven't been locked", func(t *testing.T) {
		assert.Fail(t, "todo")
	})

	t.Run("Remote ref equal to local current ref", func(t *testing.T) {
		assert.Fail(t, "todo")
	})

	t.Run("Local git-tracked non-ignored files had been synced to remote", func(t *testing.T) {
		assert.Fail(t, "todo")
	})
}
