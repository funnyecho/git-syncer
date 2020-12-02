package commandtest_test

import (
	"github.com/funnyecho/git-syncer/pkg/command/commandtest"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvCommandTimes(t *testing.T) {
	assert.Equal(t, 0, commandtest.EnvUseCommandTimes())

	assert.Equal(t, "__GO_COMMAND_TIMES__=1", commandtest.EnvWithCommandTimes(1))

	os.Setenv("__GO_COMMAND_TIMES__", "1")
	assert.Equal(t, 1, commandtest.EnvUseCommandTimes())

	assert.Equal(t, "__GO_COMMAND_TIMES__=100", commandtest.EnvWithCommandTimes(100))
	os.Setenv("__GO_COMMAND_TIMES__", "100")
	assert.Equal(t, 100, commandtest.EnvUseCommandTimes())
}

func TestEnvGoWantHelperProcess(t *testing.T) {
	assert.Equal(t, "__GO_WANT_HELPER_PROCESS__=1", commandtest.EnvWithGoWantHelperProcess())

	os.Setenv("__GO_WANT_HELPER_PROCESS__", "0")
	assert.False(t, commandtest.EnvUseGoWantHelperProcess())

	os.Setenv("__GO_WANT_HELPER_PROCESS__", "1")
	assert.True(t, commandtest.EnvUseGoWantHelperProcess())

	os.Setenv("__GO_WANT_HELPER_PROCESS__", "2")
	assert.False(t, commandtest.EnvUseGoWantHelperProcess())
}
