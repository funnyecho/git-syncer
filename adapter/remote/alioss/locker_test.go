package alioss

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/funnyecho/git-syncer/adapter/remote/alioss/buckettest"
	"github.com/funnyecho/git-syncer/adapter/remote/alioss/optionstest"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestLock(t *testing.T) {
	t.Run("failed to get remote lock info", func(t *testing.T) {
		ao, _ := New(&optionstest.MockOptions{}, &buckettest.MockBucket{
			StubIn: buckettest.StubIn{
				IsObjectExistReturnErr: errors.Err(exitcode.RemoteInvalidLock, "mock invalid lock"),
			},
		})

		id, err := ao.lock(LockRWriter)
		assert.Equal(t, exitcode.RemoteInvalidLock, errors.GetErrorCode(err))
		assert.Empty(t, id)
	})

	t.Run("failed when remote has lock file and expire time not reach", func(t *testing.T) {
		li := newLockInfo(LockRWriter)
		marshaledLi, _ := marshalLockInfo(&li)

		ao, _ := New(&optionstest.MockOptions{}, &buckettest.MockBucket{
			StubIn: buckettest.StubIn{
				IsObjectExistReturn: true,
				GetObjectReturn:     ioutil.NopCloser(strings.NewReader(string(marshaledLi))),
			},
		})

		id, err := ao.lock(LockRWriter)
		assert.Equal(t, exitcode.RemoteLocked, errors.GetErrorCode(err))
		assert.Empty(t, id)
	})

	t.Run("completed if remote has lock file but expire time had reached", func(t *testing.T) {
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
