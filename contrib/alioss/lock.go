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

/*
Lock file format:

Mutex: RLock // RLock, RWLock
Date: Mon Jan 18 16:41:44 2021 +0800 	// lock date
Locker: SamHwang1990 <samhwang1990@gmail.com>
LockID: [uuid.v4]
WLockSHA1: 664c7be795e0dce15586207234bdb2ab0d7da844 // lock writer to target sha1 (local repo sha1, not synced sha1)

*/

type lockType int8

// LockInfo lock info
type LockInfo struct {
	Mutex     lockType `json:"Mutex"`
	Date      JSONTime `json:"Date"`
	Locker    string   `json:"Locker"`
	LockID    string   `json:"LockID"`
	WLockSHA1 string   `json:"WLockSHA1"`
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

func (a *Alioss) lock(info LockInfo) (string, error) {
	if locked, lockCheckErr := a.isObjectExisted(ObjectLockFile); lockCheckErr != nil {
		return "", errors.NewError(
			errors.WithErr(lockCheckErr),
			errors.WithMsg("failed to check whether contrib was locked"),
		)
	} else if locked {
		return "", errors.NewError(
			errors.WithMsg("contrib was locked already"),
			errors.WithCode(exitcode.ContribLocked),
		)
	}

	if info.LockID == "" {
		(&info).LockID = uuid.New().String()
	}

	jsonInfo, jsonErr := json.Marshal(info)
	if jsonErr != nil {
		return "", errors.NewError(
			errors.WithErr(jsonErr),
			errors.WithMsgf("failed to json marshal lock info: %v", info),
			errors.WithCode(exitcode.ContribInvalidLock),
		)
	}

	_, uploadErr := a.uploadObject(ObjectLockFile, bytes.NewBuffer(jsonInfo), oss.ACL(oss.ACLPrivate))
	if uploadErr != nil {
		return "", errors.NewError(
			errors.WithErr(uploadErr),
			errors.WithMsgf("failed to upload lock file"),
			errors.WithCode(exitcode.ContribSyncFailed),
		)
	}

	return info.LockID, nil
}

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

	lockFileReader, lockFileReaderErr := a.getObject(ObjectLockFile)
	if lockFileReaderErr != nil {
		return errors.NewError(
			errors.WithErr(lockFileReaderErr),
			errors.WithMsg("failed to download lock file"),
			errors.WithCode(exitcode.ContribForbidden),
		)
	}

	infoBuf := new(bytes.Buffer)
	_, infoReadErr := infoBuf.ReadFrom(lockFileReader)
	if infoReadErr != nil {
		return errors.NewError(
			errors.WithErr(lockFileReaderErr),
			errors.WithMsg("failed to read contrib lock file"),
			errors.WithCode(exitcode.ContribForbidden),
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
			errors.WithCode(exitcode.ContribForbidden),
		)
	}

	return nil
}

func (a *Alioss) getLockInfo(lt lockType) LockInfo {
	return LockInfo{
		Mutex:  lt,
		Date:   JSONTime{time.Now()},
		Locker: a.getExecutor(),
		LockID: uuid.New().String(),
		// WLockSHA1: wLockSHA1,
	}
}

func (a *Alioss) getLockInfoWithWLockSHA1(lt lockType, sha1 string) LockInfo {
	info := a.getLockInfo(lt)

	if lt&LockWriter != LockIdle {
		info.WLockSHA1 = sha1
	}

	return info
}
