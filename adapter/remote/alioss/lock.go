package alioss

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/google/uuid"
)

type lockType int8

// LockInfo lock info
type LockInfo struct {
	Mutex  lockType `json:"Mutex"`
	Date   JSONTime `json:"Date"`
	LockID string   `json:"LockID"`
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

func (a *Alioss) lock(lt lockType) (string, error) {
	if locked, lockCheckErr := a.isObjectExisted(ObjectLockFile); lockCheckErr != nil {
		return "", errors.NewError(
			errors.WithErr(lockCheckErr),
			errors.WithMsg("failed to check whether contrib was locked"),
		)
	} else if locked {
		return "", errors.NewError(
			errors.WithMsg("contrib was locked already"),
			errors.WithCode(exitcode.RemoteLocked),
		)
	}

	info := a.getLockInfo(lt)

	jsonInfo, jsonErr := json.Marshal(info)
	if jsonErr != nil {
		return "", errors.NewError(
			errors.WithErr(jsonErr),
			errors.WithMsgf("failed to json marshal lock info: %v", info),
			errors.WithCode(exitcode.RemoteInvalidLock),
		)
	}

	_, uploadErr := a.uploadObject(ObjectLockFile, bytes.NewBuffer(jsonInfo), oss.ObjectACL(oss.ACLPrivate))
	if uploadErr != nil {
		return "", errors.NewError(
			errors.WithErr(uploadErr),
			errors.WithMsgf("failed to upload lock file"),
			errors.WithCode(exitcode.RemoteSyncFailed),
		)
	}

	return info.LockID, nil
}

func (a *Alioss) unlock(id string) error {
	if locked, lockCheckErr := a.isObjectExisted(ObjectLockFile); lockCheckErr != nil {
		return errors.NewError(
			errors.WithErr(lockCheckErr),
			errors.WithMsg("failed to check whether contrib was locked"),
			errors.WithCode(exitcode.RemoteForbidden),
		)
	} else if !locked {
		return errors.NewError(
			errors.WithMsg("contrib is unlock"),
			errors.WithCode(exitcode.RemoteUnlock),
		)
	}

	lockFileReader, lockFileReaderErr := a.getObject(ObjectLockFile)
	if lockFileReaderErr != nil {
		return errors.NewError(
			errors.WithErr(lockFileReaderErr),
			errors.WithMsg("failed to download lock file"),
			errors.WithCode(exitcode.RemoteForbidden),
		)
	}

	infoBuf := new(bytes.Buffer)
	_, infoReadErr := infoBuf.ReadFrom(lockFileReader)
	if infoReadErr != nil {
		return errors.NewError(
			errors.WithErr(lockFileReaderErr),
			errors.WithMsg("failed to read contrib lock file"),
			errors.WithCode(exitcode.RemoteForbidden),
		)
	}

	var info LockInfo
	jsonErr := json.Unmarshal(infoBuf.Bytes(), &info)
	if jsonErr != nil {
		return errors.NewError(
			errors.WithErr(jsonErr),
			errors.WithMsgf("failed to unmarshal contrib lock file: %s", infoBuf.String()),
			errors.WithCode(exitcode.Unknown),
		)
	}

	if info.LockID != id {
		return errors.NewError(
			errors.WithMsgf("failed to unlock contrib with different lockID; expect %s, got %s", info.LockID, id),
			errors.WithCode(exitcode.Unknown),
		)
	}

	_, unLockErr := a.deleteObject(ObjectLockFile)
	if unLockErr != nil {
		return errors.NewError(
			errors.WithErr(jsonErr),
			errors.WithMsgf("failed to delete contrib lock file"),
			errors.WithCode(exitcode.RemoteForbidden),
		)
	}

	return nil
}

func (a *Alioss) getLockInfo(lt lockType) LockInfo {
	return LockInfo{
		Mutex:  lt,
		Date:   JSONTime{time.Now()},
		LockID: uuid.New().String(),
	}
}
