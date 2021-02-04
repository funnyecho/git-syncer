package alioss

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/funnyecho/git-syncer/adapter/contrib/alioss/buckettest"
	"github.com/funnyecho/git-syncer/adapter/contrib/alioss/optionstest"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestLock(t *testing.T) {
	t.Run("failed to get contrib lock info", func(t *testing.T) {
		ao, _ := New(&optionstest.MockOptions{}, &buckettest.MockBucket{
			StubIn: buckettest.StubIn{
				IsObjectExistReturnErr: errors.Err(exitcode.ContribInvalidLock, "mock invalid lock"),
			},
		})

		id, err := ao.lock(LockRWriter)
		assert.Equal(t, exitcode.ContribInvalidLock, errors.GetErrorCode(err))
		assert.Empty(t, id)
	})

	t.Run("failed when contrib has lock file and expire time not reach", func(t *testing.T) {
		li := newLockInfo(LockRWriter)
		marshaledLi, _ := marshalLockInfo(&li)

		ao, _ := New(&optionstest.MockOptions{}, &buckettest.MockBucket{
			StubIn: buckettest.StubIn{
				IsObjectExistReturn: true,
				GetObjectReturn:     ioutil.NopCloser(strings.NewReader(string(marshaledLi))),
			},
		})

		id, err := ao.lock(LockRWriter)
		assert.Equal(t, exitcode.ContribLocked, errors.GetErrorCode(err))
		assert.Empty(t, id)
	})

	t.Run("completed if contrib has lock file but expire time had reached", func(t *testing.T) {
		li := newLockInfo(LockRWriter)
		li.ExpireDate = JSONTime{time.Now().Add(time.Second * 2)}
		marshaledLi, _ := marshalLockInfo(&li)

		ao, _ := New(&optionstest.MockOptions{}, &buckettest.MockBucket{
			StubIn: buckettest.StubIn{
				IsObjectExistReturn: true,
				GetObjectReturn:     ioutil.NopCloser(strings.NewReader(string(marshaledLi))),
			},
		})

		time.Sleep(time.Second * 4)

		id, err := ao.lock(LockRWriter)
		assert.Nil(t, err)
		assert.NotEmpty(t, id)
	})
}
