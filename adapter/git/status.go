package git

import (
	"bufio"
	"strings"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

func checkDirtyRepository() error {
	status, err := getPorcelainStatus(statusOptions{true})
	if err != nil {
		return errors.WrapC(err, exitcode.RepoUnknown, "failed to get uno porcelain status")
	}

	if len(status) > 0 {
		return errors.Err(exitcode.RepoDirty, "repository is dirty, commit changed files first")
	}

	return nil
}

// statusOptions git status options
type statusOptions struct {
	Uno bool
}

func getPorcelainStatus(options statusOptions) (status []string, err error) {
	args := []string{"status", "--porcelain"}

	if options.Uno {
		args = append(args, "-uno")
	}

	v, err := output(args)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(v))
	for scanner.Scan() {
		status = append(status, scanner.Text())
	}

	return
}
