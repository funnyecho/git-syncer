package alioss

import (
	"io"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

// getObject get object in bucket
func (a *Alioss) getObject(path string) (io.ReadCloser, error) {
	bucket, bucketErr := a.client.Bucket(a.options.Bucket)
	if bucketErr != nil {
		return nil, bucketErr
	}

	return bucket.GetObject(a.pathToKey(path), nil)
}

// uploadObject upload object
func (a *Alioss) uploadObject(path string, stream io.Reader) (string, error) {
	bucket, bucketErr := a.client.Bucket(a.options.Bucket)
	if bucketErr != nil {
		return "", bucketErr
	}

	key := a.pathToKey(path)
	uploadErr := bucket.PutObject(
		key,
		stream,
	)

	if uploadErr != nil {
		return "", errors.NewError(
			errors.WithErr(uploadErr),
			errors.WithMsgf("failed to upload file %s to bucket %s", path, bucket.BucketName),
		)
	}

	return key, nil
}

// deleteObject delete object
func (a *Alioss) deleteObject(path string) (string, error) {
	bucket, bucketErr := a.client.Bucket(a.options.Bucket)
	if bucketErr != nil {
		return "", bucketErr
	}

	key := a.pathToKey(path)

	deleteErr := bucket.DeleteObject(key)
	if deleteErr != nil {
		return "", deleteErr
	}

	return key, nil
}

func (a *Alioss) isObjectExisted(path string) (bool, error) {
	if path == "" {
		return false, errors.NewError(errors.WithMsg("path is required to check whether object exited"), errors.WithCode(exitcode.MissingArguments))
	}

	bucket, bucketErr := a.client.Bucket(a.options.Bucket)
	if bucketErr != nil {
		return false, bucketErr
	}

	return bucket.IsObjectExist(a.pathToKey(path))
}

func (a *Alioss) putSymlink(srcPath, linkPath string) error {
	if srcPath == "" || linkPath == "" {
		return errors.NewError(errors.WithMsg("srcPath or linkPath is required to put symlink"), errors.WithCode(exitcode.MissingArguments))
	}

	bucket, bucketErr := a.client.Bucket(a.options.Bucket)
	if bucketErr != nil {
		return bucketErr
	}

	return bucket.PutSymlink(a.pathToKey(linkPath), a.pathToKey(srcPath))
}
