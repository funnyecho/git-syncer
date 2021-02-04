package alioss

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/google/uuid"
)

type lockType int8

// LockInfo lock info
type LockInfo struct {
	Mutex      lockType `json:"Mutex"`
	Date       JSONTime `json:"Date"`
	ExpireDate JSONTime `json:"ExpireDate"`
	LockID     string   `json:"LockID"`
}

const (
	// LockIdle unlock
	LockIdle lockType = 0b0000
	// LockReader reader lock type
	LockReader lockType = 0b0001
	// LockWriter writer lock type
	LockWriter lockType = 0b0010
	// LockRWriter reader and writer lock type
	LockRWriter lockType = LockReader | LockWriter
)

// LockMaxAge lock max age
const LockMaxAge = time.Minute * 5

// lock contrib with lock type
func (a *Alioss) lock(lt lockType) (string, error) {
	if locked, lockCheckErr := a.isObjectExisted(ObjectLockFile); lockCheckErr != nil {
		return "", errors.NewError(
			errors.WithErr(lockCheckErr),
			errors.WithMsg("failed to check whether contrib was locked"),
		)
	} else if locked {
		info, fetchErr := a.fetchLockInfo()
		if fetchErr != nil {
			return "", fetchErr
		}

		if info.ExpireDate.After(time.Now()) {
			return "", errors.NewError(
				errors.WithMsgf("contrib was locked already, will be expired at %v", info.ExpireDate),
				errors.WithCode(exitcode.ContribLocked),
			)
		}
	}

	info := newLockInfo(lt)

	infoBytes, marshalErr := marshalLockInfo(&(info))
	if marshalErr != nil {
		return "", marshalErr
	}

	log.Debugw("lock alioss", "lockInfo", info)

	_, uploadErr := a.uploadObject(ObjectLockFile, bytes.NewBuffer(infoBytes), oss.ObjectACL(oss.ACLPrivate))
	if uploadErr != nil {
		return "", errors.NewError(
			errors.WithErr(uploadErr),
			errors.WithMsgf("failed to upload lock file"),
			errors.WithCode(exitcode.ContribSyncFailed),
		)
	}

	return info.LockID, nil
}

// unlock contrib with id
func (a *Alioss) unlock(id string) error {
	if locked, lockCheckErr := a.isObjectExisted(ObjectLockFile); lockCheckErr != nil {
		return errors.NewError(
			errors.WithErr(lockCheckErr),
			errors.WithMsg("failed to check whether contrib was locked"),
			errors.WithCode(exitcode.ContribForbidden),
		)
	} else if !locked {
		return errors.NewError(
			errors.WithMsg("contrib is unlock"),
			errors.WithCode(exitcode.ContribUnlock),
		)
	}

	info, fetchErr := a.fetchLockInfo()
	if fetchErr != nil {
		return fetchErr
	}

	if info.LockID != id {
		return errors.NewError(
			errors.WithMsgf("failed to unlock contrib with different lockID; expect %s, got %s", info.LockID, id),
			errors.WithCode(exitcode.Unknown),
		)
	}

	log.Debugw("unlock alioss", "exitedLockInfo", info)

	_, unLockErr := a.deleteObject(ObjectLockFile)
	if unLockErr != nil {
		return errors.NewError(
			errors.WithErr(unLockErr),
			errors.WithMsgf("failed to delete contrib lock file"),
			errors.WithCode(exitcode.ContribForbidden),
		)
	}

	return nil
}

func (a *Alioss) fetchLockInfo() (*LockInfo, error) {
	lockFileReader, lockFileReaderErr := a.getObject(ObjectLockFile)
	if lockFileReaderErr != nil {
		return nil, errors.NewError(
			errors.WithErr(lockFileReaderErr),
			errors.WithMsg("failed to download lock file"),
			errors.WithCode(exitcode.ContribForbidden),
		)
	}

	infoBuf := new(bytes.Buffer)
	_, infoReadErr := infoBuf.ReadFrom(lockFileReader)
	if infoReadErr != nil {
		return nil, errors.NewError(
			errors.WithErr(lockFileReaderErr),
			errors.WithMsg("failed to read contrib lock file"),
			errors.WithCode(exitcode.ContribForbidden),
		)
	}

	var info LockInfo
	if unmarshalErr := unmarshalLockInfo(&info, infoBuf.Bytes()); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	log.Debugw("fetch lock info from alioss", "lockInfo", info)

	return &info, nil
}

func newLockInfo(lt lockType) LockInfo {
	return LockInfo{
		Mutex:      lt,
		Date:       JSONTime{time.Now()},
		ExpireDate: JSONTime{time.Now().Add(LockMaxAge)},
		LockID:     uuid.New().String(),
	}
}

func unmarshalLockInfo(i *LockInfo, b []byte) error {
	jsonErr := json.Unmarshal(b, i)

	if jsonErr != nil {
		return errors.NewError(
			errors.WithErr(jsonErr),
			errors.WithMsgf("failed to unmarshal lock info: %s", string(b)),
			errors.WithCode(exitcode.Unknown),
		)
	}

	return nil
}
func marshalLockInfo(i *LockInfo) ([]byte, error) {
	jsonInfo, jsonErr := json.Marshal(i)
	if jsonErr != nil {
		return nil, errors.NewError(
			errors.WithErr(jsonErr),
			errors.WithMsgf("failed to json marshal lock info: %v", i),
			errors.WithCode(exitcode.ContribInvalidLock),
		)
	}

	return jsonInfo, nil
}
