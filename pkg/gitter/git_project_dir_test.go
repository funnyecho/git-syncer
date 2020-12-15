package gitter_test

import (
	"github.com/funnyecho/git-syncer/pkg/command"
	"github.com/funnyecho/git-syncer/pkg/command/commandtest"
	"github.com/funnyecho/git-syncer/pkg/gitter"
	. "github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestGitProjectDirCommandProcess(t *testing.T) {
	commandtest.TestHelperProcess(
		t,
		commandtest.CallTimesTestHelperProcess(
			func(name string, args ...string) (statusCode int, output string) {
				return 2, "not a git repository"
			},
			func(name string, args ...string) (statusCode int, output string) {
				return 0, "foo/bar"
			},
		),
	)
}

func TestGetProjectDir(t *testing.T) {
	mockCommander := commandtest.NewTestingCommander(
		commandtest.WithCommanderHelperName("TestGitProjectDirCommandProcess"),
	)

	oriCommander := command.PopCommand()
	defer command.PushCommand(oriCommander)

	command.PushCommand(mockCommander)

	dir, err := gitter.GetProjectDir()

	NotNil(t, err)
	if exitErr, isExitErr := err.(*exec.ExitError); isExitErr {
		Equal(t, "not a git repository", string(exitErr.Stderr))
	} else {
		Fail(t, "shall be an ExitError")
	}
	Equal(t, "exit status 2", err.Error())
	Empty(t, dir)

	dir, err = gitter.GetProjectDir()
	Nil(t, err)
	Equal(t, "foo/bar", dir)
}
