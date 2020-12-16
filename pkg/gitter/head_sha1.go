package gitter

import "github.com/funnyecho/git-syncer/pkg/command"

func GetHeadSHA1() (string, error) {
	cmd := command.Command("git", "log", "-n 1", "--pretty=format:%H")

	v, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(v), nil
}
