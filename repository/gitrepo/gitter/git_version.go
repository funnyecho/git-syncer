package gitter

import (
	"strconv"
	"strings"
)

func (g *git) GetVersion() (majorVersion, minorVersion int, err error) {
	majorVersion = 0
	minorVersion = 0

	cmd := g.command("git", "--version")

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
