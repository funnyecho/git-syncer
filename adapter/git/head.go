package git

import (
	"strings"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

// GetHead get repo head, symbolic -> sha1
func (g *Git) GetHead() (string, error) {
	var head string

	if headSymbolic, headSymbolicErr := getHeadSymbolicRef(); headSymbolicErr == nil && headSymbolic != "" {
		head = headSymbolic
	} else if headRev, headRevErr := getHeadRevision(); headRevErr == nil && headRev != "" {
		head = headRev
	} else if headSHA1, headSHA1Err := getHeadSHA1(); headSHA1Err == nil {
		head = headSHA1
	}

	if head == "" {
		return "", errors.Err(exitcode.RepoHeadNotFound, "faild to get repo head")
	}

	return head, nil
}

// GetHeadSHA1 get repo head sha1
func (g *Git) GetHeadSHA1() (string, error) {
	return getHeadSHA1()
}

// PushHead checkout to head and return original head
func (g *Git) PushHead(head string) (string, error) {
	oriHead, oriHeadErr := g.GetHead()
	if oriHeadErr != nil {
		return "", errors.Wrap(oriHeadErr, "failed to get head before checkout")
	}

	checkoutErr := checkout(head)
	if checkoutErr != nil {
		return "", errors.Wrap(checkoutErr, "failed to checkout to head %s", head)
	}

	return oriHead, nil
}

// PopHead checkout to head
func (g *Git) PopHead(head string) error {
	return checkout(head)
}

func getHeadSHA1() (string, error) {
	return output([]string{"log", "-n 1", "--pretty=format:%H"})
}

func getHeadRevision() (string, error) {
	return output([]string{"rev-parse", "HEAD"})
}

func getHeadSymbolicRef() (string, error) {
	v, err := output([]string{"symbolic-ref", "HEAD"})

	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(v, "refs/heads/"), nil
}

func checkout(head string) error {
	if head == "" {
		return errors.Err(exitcode.InvalidParams, "missing head to checkout")
	}

	_, err := output([]string{"checkout", head})
	return err
}
