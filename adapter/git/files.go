package git

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

// ListTrackedFiles list all tracked files
func (g *Git) ListTrackedFiles(path string) ([]string, error) {
	if gitVerErr := checkVersion(); gitVerErr != nil {
		return nil, gitVerErr
	}

	if dirtyErr := checkDirtyRepository(); dirtyErr != nil {
		return nil, dirtyErr
	}

	return listTrackedFiles(path)
}

// ListChangedFiles list changed files between baseSha1 and head
func (g *Git) ListChangedFiles(path string, baseSha1 string) (amFiles []string, dFiles []string, err error) {
	if gitVerErr := checkVersion(); gitVerErr != nil {
		return nil, nil, gitVerErr
	}

	if dirtyErr := checkDirtyRepository(); dirtyErr != nil {
		return nil, nil, dirtyErr
	}

	defer func() {
		if err != nil {
			amFiles = nil
			dFiles = nil
		}
	}()

	amFiles, amErr := diff(path, baseSha1, diffOptions{"AM"})
	if amErr != nil {
		err = errors.WrapC(amErr, exitcode.RepoListFilesFailed, "failed to list files with filter `AM`")
		return
	}

	dFiles, dErr := diff(path, baseSha1, diffOptions{"D"})
	if dErr != nil {
		err = errors.WrapC(dErr, exitcode.RepoListFilesFailed, "failed to list files with filter `D`")
		return
	}

	return
}

func listTrackedFiles(path string) (files []string, err error) {
	if path == "" {
		path = "."
	}

	v, err := output([]string{"ls-files", "--", path})
	if err != nil {
		return nil, err
	}

	var fs []string

	scanner := bufio.NewScanner(strings.NewReader(v))
	for scanner.Scan() {
		fs = append(fs, scanner.Text())
	}

	return fs, nil
}

type diffOptions struct {
	filter string
}

func diff(path string, commit string, options diffOptions) ([]string, error) {
	args := []string{"diff", "--name-only", "--no-renames"}

	if options.filter != "" {
		args = append(args, fmt.Sprintf("--diff-filter=%s", options.filter))
	}

	if path == "" {
		path = "."
	}

	args = append(args, commit, "--", path)

	var files []string

	v, err := output(args)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(v))
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	return files, err
}
