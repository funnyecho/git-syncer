package gitrepo

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

func (r *repo) ListAllFiles() (sha1 string, uploads []string, err error) {
	if gitVErr := r.validateGitVersion(); gitVErr != nil {
		return "", nil, gitVErr
	}

	if isDirtyRepo, repoStatusErr := IsDirtyRepository(r.gitter); repoStatusErr != nil {
		return "", nil, errors.NewError(
			errors.WithMsg("failed to get check repository status"),
			errors.WithErr(repoStatusErr),
		)
	} else if isDirtyRepo {
		return "", nil, errors.NewError(
			errors.WithCode(exitcode.RepoDirty),
			errors.WithMsg("dirty repository: Having uncommitted changes"),
		)
	}

	sha1, sha1Err := r.GetHeadSHA1()
	if sha1Err != nil {
		return "", nil, sha1Err
	}

	files, filesErr := r.gitter.ListFiles(GetSyncRoot(r))
	if filesErr != nil {
		return "", nil, filesErr
	}

	return sha1, files, nil
}

func (r *repo) ListChangedFiles(baseSha1 string) (sha1 string, uploads []string, deletes []string, err error) {
	if baseSha1 == "" {
		return "", nil, nil, errors.NewError(
			errors.WithCode(exitcode.RepoDiffBaseNotFound),
			errors.WithMsg("basic sha1 to diff changed files is empty"),
		)
	}

	if gitVErr := r.validateGitVersion(); gitVErr != nil {
		return "", nil, nil, gitVErr
	}

	if isDirtyRepo, repoStatusErr := IsDirtyRepository(r.gitter); repoStatusErr != nil {
		return "", nil, nil, errors.NewError(
			errors.WithMsg("failed to get check repository status"),
			errors.WithErr(repoStatusErr),
		)
	} else if isDirtyRepo {
		return "", nil, nil, errors.NewError(
			errors.WithCode(exitcode.RepoDirty),
			errors.WithMsg("dirty repository: Having uncommitted changes"),
		)
	}

	repoSHA1, repoSHA1Err := r.gitter.GetHeadSHA1()
	if repoSHA1Err != nil {
		err = errors.NewError(
			errors.WithCode(exitcode.RepoUnknown),
			errors.WithMsg("failed to get repo head sha1"),
			errors.WithErr(repoSHA1Err),
		)
		return
	} else if repoSHA1 == "" {
		err = errors.NewError(
			errors.WithMsg("repo head sha1 not found"),
			errors.WithCode(exitcode.RepoHeadNotFound),
		)
		return
	}

	syncRoot := GetSyncRoot(r)

	diffAM, diffAMErr := r.gitter.DiffAM(syncRoot, baseSha1)
	if diffAMErr != nil {
		err = errors.NewError(
			errors.WithCode(exitcode.RepoUnknown),
			errors.WithMsgf("failed to diff append or modified files from %s in path %s", baseSha1, GetSyncRoot(r)),
			errors.WithErr(diffAMErr),
		)
		return
	}

	diffD, diffAMErr := r.gitter.DiffD(syncRoot, baseSha1)
	if diffAMErr != nil {
		err = errors.NewError(
			errors.WithCode(exitcode.RepoUnknown),
			errors.WithMsgf("failed to diff deleted files from %s in path %s", baseSha1, GetSyncRoot(r)),
			errors.WithErr(diffAMErr),
		)
		return
	}

	return repoSHA1, diffAM, diffD, nil
}

func (r *repo) validateGitVersion() error {
	if gitMajorV, gitMinorV, gitVerErr := r.gitter.GetVersion(); gitVerErr != nil {
		return errors.NewError(errors.WithMsg("Git haven't installed? "), errors.WithErr(gitVerErr))
	} else if gitMajorV < 1 || (gitMajorV < 2 && gitMinorV < 7) {
		return errors.NewError(errors.WithMsg("git is too old, 1.7.0 or higher supported only"))
	}

	return nil
}
