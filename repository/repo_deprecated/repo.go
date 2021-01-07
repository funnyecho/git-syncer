package repo

import (
	"fmt"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/fs"
	"github.com/funnyecho/git-syncer/pkg/gitter"
	"github.com/funnyecho/git-syncer/repository"
)

type Options struct {
	WorkingDir string
}

func New(options Options) *repo {
	return &repo{
		Options: &options,
	}
}

// Interface implement checking
var _ repository.Repository = &repo{}

type repo struct {
	*Options
}

func (r *repo) GetConfig(keys ...string) (string, error) {
	panic("implement me")
}

func (r *repo) GetHead() (string, error) {
	var head string

	if symbolicHead := gitter.GetSymbolicHead(); symbolicHead != "" {
		head = symbolicHead
	} else if headRev := gitter.GetHead(); headRev != "" {
		head = headRev
	} else if localSHA1, localSHA1Err := gitter.GetHeadSHA1(); localSHA1Err == nil {
		head = localSHA1
	}

	return head, nil
}

func (r *repo) GetHeadSHA1() (string, error) {
	return gitter.GetHeadSHA1()
}

func (r *repo) GetRepoDir() (string, error) {
	return gitter.GetProjectDir()
}

func (r *repo) GetSyncRoot() (string, error) {
	root, _ := r.GetConfig("sync_root")
	if root == "" {
		root = "./assets"
	}

	if rootExisted, rootErr := fs.IsDirExists(root); rootExisted {
		return root, nil
	} else {
		return "", errors.NewError(errors.WithMsg("sync root not an valid directory"), errors.WithErr(rootErr))
	}
}

func (r *repo) GitVersion() (majorV, minorV int, err error) {
	return gitter.GetGitVersion()
}

func (r *repo) IsDirtyRepository() (bool, error) {
	status, err := gitter.GetPorcelainStatus(gitter.WithUnoPorcelainStatus())
	if err != nil {
		return false, err
	}

	if len(status) > 0 {
		return false, nil
	}

	return true, nil
}

func (r *repo) ListAllFiles(syncRoot string) ([]string, error) {
	return gitter.ListFiles(syncRoot)
}

func (r *repo) ListChangedFiles(syncRoot string) (upload []string, delete []string, err error) {
	return
}

func (r *repo) SetHead(head string) error {
	return gitter.Checkout(head)
}

func (r *repo) SetupTempDir() (string, error) {
	tempDirFallbackOptions := []fs.TempDirFallbackOption{
		fs.WithTempDirFallback("", "git-syncer-*"),
	}

	if repoDir, repoDirErr := gitter.GetProjectDir(); repoDirErr == nil {
		tempDirFallbackOptions = append(
			tempDirFallbackOptions,
			fs.WithTempDirFallback(fmt.Sprintf("%s/.git", repoDir), "git-syncer-temp"),
		)
	}

	dir, errStack := fs.CreateTempDir(tempDirFallbackOptions...)
	if errStack != nil {
		return "", errors.NewError(errors.WithMsg(fmt.Sprintf("errors: %v", errStack)))
	}

	return dir, nil
}
