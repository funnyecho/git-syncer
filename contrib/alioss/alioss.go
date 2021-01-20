package alioss

import (
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/contrib"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/fs"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/repository"
)

// NewContribFactory create alioss contrib factory
func NewContribFactory() contrib.Factory {
	return func(options interface{ repository.ConfigReader }) (contrib.Contrib, error) {
		configurable := contrib.NewConfigurable("alioss", options)

		ossOptions, ossOptionsErr := NewOptions(configurable)
		if ossOptionsErr != nil {
			return nil, ossOptionsErr
		}

		ossClient, ossClientErr := oss.New(ossOptions.Endpoint, ossOptions.AccessKeyID, ossOptions.AccessKeySecret)
		if ossClientErr != nil {
			return nil, ossClientErr
		}

		c := &Alioss{
			configurable,
			ossOptions,
			ossClient,
		}

		return c, nil
	}
}

// Alioss alioss contrib
type Alioss struct {
	*contrib.Configurable
	options *Options

	client *oss.Client
}

// GetHeadSHA1 get head sha1 from alioss contrib
func (a *Alioss) GetHeadSHA1() (string, error) {
	rLockID, rLockErr := a.lock(a.getLockInfo(LockReader))
	if rLockErr != nil {
		return "", rLockErr
	}

	defer func() {
		unLockErr := a.unlock(rLockID)
		if unLockErr != nil {
			log.Errore("failed to unlock reader", unLockErr, "lockID", rLockID)
		}
	}()

	head, headErr := a.peekLog()
	if headErr != nil {
		return "", headErr
	}

	return head.SHA1, nil
}

// Sync sync files to alioss contrib
func (a *Alioss) Sync(reqx *contrib.SyncReq) (res contrib.SyncRes, err error) {
	if reqx == nil {
		return contrib.SyncRes{}, errors.NewError(errors.WithCode(exitcode.MissingArguments), errors.WithMsg("failed to sync without request info"))
	}

	lockID, lockErr := a.lock(a.getLockInfoWithWLockSHA1(LockRWriter, reqx.SHA1))
	if lockErr != nil {
		return contrib.SyncRes{}, lockErr
	}
	defer func() {
		unlockErr := a.unlock(lockID)
		if unlockErr != nil {
			log.Errore("failed to unlock writer", unlockErr, "lockID", lockID)
		}
	}()

	var uploaded []UploadedFile
	var deleted []string

	defer func() {
		var uploadedPaths []string

		for _, uf := range uploaded {
			uploadedPaths = append(uploadedPaths, uf.Path)
		}

		res = contrib.SyncRes{
			Uploaded: uploadedPaths,
			Deleted:  deleted,
		}

		if err == nil {
			res.SHA1 = reqx.SHA1
		}
	}()

	for _, p := range reqx.Uploads {
		f, fErr := os.Open(p)
		if fErr != nil {
			err = fErr
			return
		}
		defer f.Close()

		_, uErr := a.uploadObject(p, f)
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
			udF.SHA1 = string(fSHA1)
		}

		uploaded = append(uploaded, udF)
	}

	for _, p := range reqx.Deletes {
		_, uErr := a.deleteObject(p)
		if uErr != nil {
			err = uErr
			return
		}

		deleted = append(deleted, p)
	}

	head, headErr := a.peekLog()
	if headErr != nil {
		log.Infow("failed to get head log", "err", headErr)
	} else if head == nil {
		log.Infow("contrib head log is empty")
	}

	logInfo := LogInfo{
		SHA1:     reqx.SHA1,
		RefSHA1:  head.SHA1,
		Executor: a.getExecutor(),
		Date:     JSONTime(time.Now()),
		Uploaded: uploaded,
		Deleted:  deleted,
	}
	if logErr := a.pushLog(logInfo); logErr != nil {
		err = logErr
		return
	}

	return
}
