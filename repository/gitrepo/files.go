package gitrepo

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

func (r *repo) ListAllFiles() (sha1 string, uploads []string, err error) {
	if gitVErr := r.validateGitVersion(); gitVErr != nil {
		return "", nil, gitVErr
	}

	if isDirtyRepo, repoStatusErr := r.IsDirtyRepository(); repoStatusErr != nil {
		return "", nil, errors.NewError(
			errors.WithStatusCode(exitcode.Git),
			errors.WithMsg("failed to get check repository status"),
			errors.WithErr(repoStatusErr),
		)
	} else if isDirtyRepo {
		return "", nil, errors.NewError(
			errors.WithStatusCode(exitcode.Git),
			errors.WithMsg("dirty repository: Having uncommitted changes"),
		)
	}

	sha1, sha1Err := r.GetHeadSHA1()
	if sha1Err != nil {
		return "", nil, sha1Err
	}

	files, filesErr := r.gitter.ListFiles(r.syncRoot)
	if filesErr != nil {
		return "", nil, filesErr
	}

	return sha1, files, nil
}

func (r *repo) ListChangedFiles(baseSha1 string) (sha1 string, uploads []string, deletes []string, err error) {
	panic("implement me")
}

func (r *repo) validateGitVersion() error {
	if gitMajorV, gitMinorV, gitVerErr := r.gitter.GetVersion(); gitVerErr != nil {
		return errors.NewError(errors.WithMsg("Git haven't installed? "), errors.WithErr(gitVerErr))
	} else if gitMajorV < 2 && gitMinorV < 7 {
		return errors.NewError(errors.WithMsg("git is too old, 1.7.0 or higher supported only"))
	}

	return nil
}
