package gitter

import (
	"github.com/funnyecho/git-syncer/pkg/command"
	"strconv"
	"strings"
)

func GetGitVersion() (majorVersion, minorVersion int, err error) {
	majorVersion = 0
	minorVersion = 0

	cmd := command.Command("git", "--version")

	plainVersion, err := cmd.Output()
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
