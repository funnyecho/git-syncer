package gitter

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func GetGitVersion() (majorVersion, minorVersion int, err error) {
	majorVersion = 0
	minorVersion = 0

	cmd := exec.Command("git", "--version")

	var stdout bytes.Buffer

	cmd.Stdout= &stdout

	err = cmd.Run()
	if err != nil {
		return
	}

	plainVersion := string(stdout.Bytes())

	version := strings.Fields(plainVersion)[2]
	versionParts := strings.FieldsFunc(version, func(r rune) bool {
		return r == '.'
	})

	majorVersion, _ = strconv.Atoi(versionParts[0])
	minorVersion, _ = strconv.Atoi(versionParts[1])
	return
}
