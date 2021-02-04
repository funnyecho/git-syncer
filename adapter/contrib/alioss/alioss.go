package alioss

import (
	"fmt"
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/funnyecho/git-syncer/pkg/fs"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/gosuri/uilive"
)

// New create alioss contrib
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

	uploadProgress := uilive.New()
	uploadProgress.Start()

	defer uploadProgress.Stop()

	fmt.Println("syncing files to alioss")
	fmt.Fprintf(uploadProgress, "Uploading.. (%d/%d) \n", 0, len(uploads))

	for i, p := range uploads {
		if udF, uploadErr := a.uploadFile(p); uploadErr != nil {
			err = uploadErr
			return
		} else {
			uploadedFiles = append(uploadedFiles, *udF)
		}

		fmt.Fprintf(uploadProgress, "Uploading.. (%d/%d) \n", i+1, len(uploads))
	}

	deleteProgress := uilive.New()
	deleteProgress.Start()

	defer deleteProgress.Stop()

	fmt.Fprintf(uploadProgress, "Deleting.. (%d/%d) \n", 0, len(deletes))

	for i, p := range deletes {
		_, uErr := a.deleteObject(p)
		if uErr != nil {
			err = uErr
			return
		}

		deleted = append(deleted, p)
		fmt.Fprintf(uploadProgress, "Deleting.. (%d/%d) \n", i+1, len(deletes))
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

func (a *Alioss) uploadFile(p string) (*UploadedFile, error) {
	f, fErr := os.Open(p)
	if fErr != nil {
		return nil, fErr
	}
	defer f.Close()

	_, uErr := a.uploadObject(p, f, oss.ObjectACL(oss.ACLPublicRead))
	if uErr != nil {
		return nil, uErr
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

	return &udF, nil
}
