package alioss

import (
	"fmt"
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/funnyecho/git-syncer/pkg/fs"
	"github.com/funnyecho/git-syncer/pkg/log"
)

// New create alioss remote
func New(opt Options, bkt Bucket) (*Alioss, error) {
	c := &Alioss{
		opt,
		bkt,
	}

	return c, nil
}

// Alioss alioss contrib
type Alioss struct {
	opts   Options
	bucket Bucket
}

// GetHeadSHA1 get head sha1 from alioss contrib
func (a *Alioss) GetHeadSHA1() (string, error) {
	rLockID, rLockErr := a.lock(LockReader)
	if rLockErr != nil {
		return "", rLockErr
	}

	defer func() {
		err := recover()

		unLockErr := a.unlock(rLockID)
		if unLockErr != nil {
			log.Errore("failed to unlock reader", unLockErr, "lockID", rLockID)
		}

		if err != nil {
			panic(err)
		}
	}()

	head, headErr := a.PeekLog()
	if headErr != nil {
		return "", headErr
	}

	return head.SHA1, nil
}

// Sync sync files to alioss contrib
func (a *Alioss) Sync(sha1 string, uploads []string, deletes []string) (uploaded []string, deleted []string, err error) {
	lockID, lockErr := a.lock(LockRWriter)
	if lockErr != nil {
		err = lockErr
		return
	}

	defer func() {
		unlockErr := a.unlock(lockID)
		if unlockErr != nil {
			log.Errore("failed to unlock writer", unlockErr, "lockID", lockID)
		}
	}()

	var uploadedFiles []UploadedFile

	defer func() {
		for _, uf := range uploadedFiles {
			uploaded = append(uploaded, uf.Path)
		}
	}()

	for _, p := range uploads {
		f, fErr := os.Open(p)
		if fErr != nil {
			err = fErr
			return
		}
		defer f.Close()

		_, uErr := a.uploadObject(p, f, oss.ObjectACL(oss.ACLPublicRead))
		if uErr != nil {
			err = uErr
			return
		}

		udF := UploadedFile{
			Path: p,
		}

		fStat, fsErr := os.Stat(p)
		if fsErr == nil {
			udF.Size = fStat.Size()
		}

		fSHA1, fSHA1Err := fs.GetFileSHA1(p)
		if fSHA1Err == nil {
			udF.SHA1 = fmt.Sprintf("%x", fSHA1)
		}

		uploadedFiles = append(uploadedFiles, udF)
	}

	for _, p := range deletes {
		_, uErr := a.deleteObject(p)
		if uErr != nil {
			err = uErr
			return
		}

		deleted = append(deleted, p)
	}

	head, headErr := a.PeekLog()
	if headErr != nil {
		log.Infow("failed to get head log", "err", headErr)
	} else if head == nil {
		log.Infow("contrib head log is empty")
	}

	logInfo := LogInfo{
		SHA1:     sha1,
		RefSHA1:  head.SHA1,
		Date:     JSONTime{time.Now()},
		Uploaded: uploadedFiles,
		Deleted:  deleted,
	}
	if logErr := a.PushLog(logInfo); logErr != nil {
		err = logErr
		return
	}

	return
}
