package gitrepo_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/funnyecho/git-syncer/repository/gitrepo"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("options fn was called", func(t *testing.T) {
		fnCalled := 0
		var fn gitrepo.WithOptions = func(o *gitrepo.Option) error {
			fnCalled++
			return nil
		}

		gitrepo.New(fn, fn, fn, fn)
		assert.Equal(t, 4, fnCalled)
	})

	t.Run("bailing options fn if error occur", func(t *testing.T) {
		fnCalled := 0
		errFnCalled := 0
		err := fmt.Errorf("something failed")

		var fn gitrepo.WithOptions = func(o *gitrepo.Option) error {
			fnCalled++
			return nil
		}

		var errFn gitrepo.WithOptions = func(o *gitrepo.Option) error {
			errFnCalled++
			return err
		}

		_, errR := gitrepo.New(fn, fn, errFn, fn, errFn)
		assert.Equal(t, err, errR)
		assert.Equal(t, 2, fnCalled)
		assert.Equal(t, 1, errFnCalled)
	})
}

func TestWithWorkingDir(t *testing.T) {

	t.Run("failed with invalid path", func(t *testing.T) {
		fn := gitrepo.WithWorkingDir("path not existed")

		err := fn(nil)
		assert.Error(t, err)
	})

	t.Run("working dir changed", func(t *testing.T) {
		wd, _ := os.Getwd()
		fn := gitrepo.WithWorkingDir(filepath.Dir(wd))

		err := fn(nil)
		assert.Nil(t, err)

		nwd, _ := os.Getwd()
		assert.Equal(t, filepath.Dir(wd), nwd)
	})

}
