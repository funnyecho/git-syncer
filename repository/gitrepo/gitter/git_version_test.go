package gitter_test

import (
	"github.com/funnyecho/git-syncer/pkg/command/commandtest"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestGitVersionCommandProcess(t *testing.T) {
	commandtest.TestHelperProcess(
		t,
		commandtest.CallTimesTestHelperProcess(
			func(name string, args ...string) (statusCode int, output string) {
				return 0, "git version 2.24.3 (Apple Git-128)"
			},
			func(name string, args ...string) (statusCode int, output string) {
				return 0, "git version 1.1.3 (Apple Git-128)"
			},
			func(name string, args ...string) (statusCode int, output string) {
				return 0, "git version 0.0.2 (Apple Git-128)"
			},
		),
	)
}

func TestGetGitVersion(t *testing.T) {
	mockCommander := commandtest.NewTestingCommander(
		commandtest.WithCommanderHelperName("TestGitVersionCommandProcess"),
	)

	git := gitter.NewDefaultGitterWithCommander(mockCommander)

	majorV, minorV, err := git.GetVersion()
	Equal(t, 2, majorV)
	Equal(t, 24, minorV)
	Nil(t, err)

	majorV, minorV, err = git.GetVersion()
	Equal(t, 1, majorV)
	Equal(t, 1, minorV)
	Nil(t, err)

	majorV, minorV, err = git.GetVersion()
	Equal(t, 0, majorV)
	Equal(t, 0, minorV)
	Nil(t, err)
}
