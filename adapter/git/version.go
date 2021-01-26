package git

import (
	"strconv"
	"strings"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

func getVersion() (majorVersion, minorVersion int, err error) {
	majorVersion = 0
	minorVersion = 0

	plainVersion, err := output([]string{"--version"})
	if err != nil {
		return
	}

	version := strings.Fields(string(plainVersion))[2]
	versionParts := strings.FieldsFunc(version, func(r rune) bool {
		return r == '.'
	})

	majorVersion, _ = strconv.Atoi(versionParts[0])
	minorVersion, _ = strconv.Atoi(versionParts[1])
	return
}

func checkVersion() error {
	if gitMajorV, gitMinorV, gitVerErr := getVersion(); gitVerErr != nil {
		return errors.WrapC(gitVerErr, exitcode.RepoUnknown, "Git haven't installed? ")
	} else if gitMajorV < 1 || (gitMajorV < 2 && gitMinorV < 7) {
		return errors.Err(exitcode.RepoInvalidGitVersion, "git is too old, 1.7.0 or higher supported only")
	}

	return nil
}
