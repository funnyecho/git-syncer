package gitrepo

import (
	"fmt"
	"os"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/repository"
	gitter2 "github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
)

// New new gitrepo instance
func New(options ...WithOptions) (repository.Repository, error) {
	o := &Option{}

	for _, fn := range options {
		if err := fn(o); err != nil {
			return nil, err
		}
	}

	gt := o.Gitter
	if gt == nil {
		gt = gitter2.NewDefaultGitter()
	}

	rp := &repo{
		gitter: gt,
		remote: o.WorkingRemote,
	}

	return rp, nil
}

type repo struct {
	remote string
	gitter gitter2.Gitter
}

// WithWorkingDir change working dir
func WithWorkingDir(dir string) WithOptions {
	return func(o *Option) error {
		if err := os.Chdir(dir); err != nil {
			return errors.NewError(
				errors.WithStatusCode(exitcode.Filesystem),
				errors.WithErr(err),
				errors.WithMsg(fmt.Sprintf("failed to change working dir to %s", dir)),
			)
		}
		return nil
	}
}

// WithWorkingRemote set working remote
func WithWorkingRemote(remote string) WithOptions {
	return func(o *Option) error {
		o.WorkingRemote = remote
		return nil
	}
}

// WithGitter set gitter implement
func WithGitter(gt gitter2.Gitter) WithOptions {
	return func(o *Option) error {
		o.Gitter = gt
		return nil
	}
}

// WithOptions func to initilize gitrepo options
type WithOptions = func(*Option) error

// Option option to init gitrepo
type Option struct {
	WorkingDir    string
	WorkingRemote string
	Gitter        gitter2.Gitter
}
