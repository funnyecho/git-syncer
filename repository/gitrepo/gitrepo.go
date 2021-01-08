package gitrepo

import (
	"fmt"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/repository"
	gitter2 "github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
	"os"
)

func New(options ...WithOptions) (*repo, error) {
	o := &option{}

	for _, fn := range options {
		fn(o)
	}

	if o.workingDir != "" {
		if err := os.Chdir(o.workingDir); err != nil {
			return nil, errors.NewError(
				errors.WithStatusCode(exitcode.Filesystem),
				errors.WithErr(err),
				errors.WithMsg(fmt.Sprintf("failed to change working dir to %s", o.workingDir)),
			)
		}
	}

	gt := o.gitter
	if gt == nil {
		gt = gitter2.NewDefaultGitter()
	}

	rp := &repo{
		gitter: gt,
		remote: o.workingRemote,
	}

	syncRootErr := rp.initSyncRoot()
	if syncRootErr != nil {
		return nil, syncRootErr
	}

	return rp, nil
}

type repo struct {
	syncRoot string
	remote   string
	gitter   gitter2.Gitter
}

var _ repository.Repository = &repo{}

func WithWorkingDir(dir string) WithOptions {
	return func(o *option) {
		o.workingDir = dir
	}
}

func WithWorkingRemote(remote string) WithOptions {
	return func(o *option) {
		o.workingRemote = remote
	}
}

func WithGitter(gt gitter2.Gitter) WithOptions {
	return func(o *option) {
		o.gitter = gt
	}
}

type WithOptions = func(*option)

type option struct {
	workingDir    string
	workingRemote string
	gitter        gitter2.Gitter
}
