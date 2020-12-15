package gitter_test

import (
	"testing"

	"github.com/funnyecho/git-syncer/pkg/command"
	"github.com/funnyecho/git-syncer/pkg/command/commandtest"
	"github.com/funnyecho/git-syncer/pkg/gitter"
	"github.com/stretchr/testify/assert"
)

func TestGitPorcelainStatusCommandProcess(t *testing.T) {
	commandtest.TestHelperProcess(
		t,
		commandtest.CallTimesTestHelperProcess(
			func(name string, args ...string) (statusCode int, output string) {
				return 2, "something error happened"
			},
			func(name string, args ...string) (statusCode int, output string) {
				return 0,
					"M .DS_Store/\n" +
						"?? .vscode/"
			},
			func(name string, args ...string) (statusCode int, output string) {
				return 0,
					"M .DS_Store/"
			},
		),
	)
}

func TestPorcelainStatus(t *testing.T) {
	mockCommander := commandtest.NewTestingCommander(
		commandtest.WithCommanderHelperName("TestGitPorcelainStatusCommandProcess"),
	)

	oriCommander := command.PopCommand()
	defer command.PushCommand(oriCommander)

	command.PushCommand(mockCommander)

	status, err := gitter.GetPorcelainStatus()

	assert.NotNil(t, err)
	assert.Equal(t, "exit status 2", err.Error())
	assert.Empty(t, status)

	status, err = gitter.GetPorcelainStatus()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(status))
	assert.Equal(t, "M .DS_Store/", status[0])
	assert.Equal(t, "?? .vscode/", status[1])

	status, err = gitter.GetPorcelainStatus(gitter.WithUnoPorcelainStatus())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(status))
	assert.Equal(t, "M .DS_Store/", status[0])
}
