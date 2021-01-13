package gitter_test

import (
	"testing"

	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter"

	"github.com/funnyecho/git-syncer/pkg/command/commandtest"
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

	git := gitter.NewDefaultGitterWithCommander(mockCommander)

	status, err := git.GetPorcelainStatus()

	assert.NotNil(t, err)
	assert.Equal(t, "exit status 2", err.Error())
	assert.Empty(t, status)

	status, err = git.GetPorcelainStatus()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(status))
	assert.Equal(t, "M .DS_Store/", status[0])
	assert.Equal(t, "?? .vscode/", status[1])

	status, err = git.GetUnoPorcelainStatus()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(status))
	assert.Equal(t, "M .DS_Store/", status[0])
}
