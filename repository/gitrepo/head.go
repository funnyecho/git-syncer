package gitrepo

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

func (r *repo) GetHead() (string, error) {
	var head string

	if symbolicHead, symbolicHeadErr := r.gitter.GetSymbolicHead(); symbolicHeadErr == nil && symbolicHead != "" {
		head = symbolicHead
	} else if headRev, headRevErr := r.gitter.GetHead(); headRevErr == nil && headRev != "" {
		head = headRev
	} else if localSHA1, localSHA1Err := r.gitter.GetHeadSHA1(); localSHA1Err == nil {
		head = localSHA1
	}

	if head == "" {
		return "", errors.NewError(errors.WithStatusCode(exitcode.Git), errors.WithMsg("failed to get repo head"))
	}

	return head, nil
}

func (r *repo) GetHeadSHA1() (string, error) {
	sha1, err := r.gitter.GetHeadSHA1()

	if err != nil {
		return "", err
	}

	if sha1 == "" {
		return "", errors.NewError(errors.WithStatusCode(exitcode.Git), errors.WithMsg("failed to get repo head sha1"))
	}

	return sha1, nil
}

func (r *repo) PushHead(head string) (string, error) {
	prevHead, prevHeadErr := r.GetHead()

	if prevHeadErr != nil {
		return "", prevHeadErr
	}

	checkoutErr := r.gitter.Checkout(head)
	if checkoutErr != nil {
		return "", checkoutErr
	}

	return prevHead, nil
}

func (r *repo) PopHead(head string) error {
	return r.gitter.Checkout(head)
}
