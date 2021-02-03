package alioss

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

// getObject get object in bucket
func (a *Alioss) getObject(path string) (io.ReadCloser, error) {
	return a.bucket.GetObject(a.PathToKey(path), nil)
}

// uploadObject upload object
func (a *Alioss) uploadObject(path string, stream io.Reader, options ...oss.Option) (string, error) {
	key := a.PathToKey(path)
	uploadErr := a.bucket.PutObject(
		key,
		stream,
		options...,
	)

	if uploadErr != nil {
		return "", errors.NewError(
			errors.WithErr(uploadErr),
			errors.WithMsgf("failed to upload file %s to bucket %s", path, a.opts.Bucket()),
		)
	}

	return key, nil
}

// deleteObject delete object
func (a *Alioss) deleteObject(path string) (string, error) {
	key := a.PathToKey(path)

	deleteErr := a.bucket.DeleteObject(key)
	if deleteErr != nil {
		return "", deleteErr
	}

	return key, nil
}

func (a *Alioss) isObjectExisted(path string) (bool, error) {
	if path == "" {
		return false, errors.NewError(errors.WithMsg("path is required to check whether object exited"), errors.WithCode(exitcode.MissingArguments))
	}

	return a.bucket.IsObjectExist(a.PathToKey(path))
}

func (a *Alioss) putSymlink(srcPath, linkPath string, options ...oss.Option) error {
	if srcPath == "" || linkPath == "" {
		return errors.NewError(errors.WithMsg("srcPath or linkPath is required to put symlink"), errors.WithCode(exitcode.MissingArguments))
	}

	return a.bucket.PutSymlink(a.PathToKey(linkPath), a.PathToKey(srcPath), options...)
}
