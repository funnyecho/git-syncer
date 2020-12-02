package commandtest_test

import (
	"github.com/funnyecho/git-syncer/pkg/command/commandtest"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func TestCallTimesTestHelperProcess(t *testing.T) {
	callTimesExecutor := commandtest.CallTimesTestHelperProcess(
		func(name string, args ...string) (statusCode int, output string) {
			return 1, "1"
		},
		func(name string, args ...string) (statusCode int, output string) {
			return 2, "2"
		},
		func(name string, args ...string) (statusCode int, output string) {
			return 3, "3"
		},
	)

	var statusCode int
	var output string

	os.Setenv(commandtest.EnvCommandTimes, "1")
	statusCode, output = callTimesExecutor("")

	assert.Equal(t, 1, statusCode)
	assert.Equal(t, "1", output)

	os.Setenv(commandtest.EnvCommandTimes, "2")
	statusCode, output = callTimesExecutor("")

	assert.Equal(t, 2, statusCode)
	assert.Equal(t, "2", output)

	os.Setenv(commandtest.EnvCommandTimes, "3")
	statusCode, output = callTimesExecutor("")

	assert.Equal(t, 3, statusCode)
	assert.Equal(t, "3", output)

	t.Run("when callTimes is large than executors length, alway call the last executor", func(t *testing.T) {
		os.Setenv(commandtest.EnvCommandTimes, "4")
		statusCode, output = callTimesExecutor("")

		assert.Equal(t, 3, statusCode)
		assert.Equal(t, "3", output)

		os.Setenv(commandtest.EnvCommandTimes, "100")
		statusCode, output = callTimesExecutor("")

		assert.Equal(t, 3, statusCode)
		assert.Equal(t, "3", output)
	})
}

func TestTestHelperProcess(t *testing.T) {
	commander := commandtest.NewTestingCommander(
		commandtest.WithCommanderHelperName("TestCommandExecutor"),
	)

	cmd := commander("foo")
	output, err := cmd.Output()

	assert.Nil(t, err)
	assert.Equal(
		t,
		"1. exit with 0\n2. exit with 0\n3.1. exit with 0\n4.1. exit with 0",
		string(output),
	)

	cmd = commander("foo")
	output, err = cmd.Output()

	assert.NotNil(t, err)
	assert.Equal(
		t,
		"1. exit with 0\n2. exit with 0",
		string(output),
	)
	if exitErr, isExitErr := err.(*exec.ExitError); isExitErr {
		assert.Equal(t, 1, exitErr.ExitCode())
		assert.Equal(t, "3.2. exit with 1", string(exitErr.Stderr))
	} else {
		assert.Fail(t, "shall be an ExitError")
	}

	cmd = commander("foo")
	output, err = cmd.Output()

	assert.NotNil(t, err)
	assert.Equal(
		t,
		"1. exit with 0\n2. exit with 0\n3.3. exit with 0",
		string(output),
	)
	if exitErr, isExitErr := err.(*exec.ExitError); isExitErr {
		assert.Equal(t, 2, exitErr.ExitCode())
		assert.Equal(t, "4.3. exit with 2", string(exitErr.Stderr))
	} else {
		assert.Fail(t, "shall be an ExitError")
	}
}

func TestCommandExecutor(t *testing.T) {
	commandtest.TestHelperProcess(
		t,
		func(name string, args ...string) (statusCode int, output string) {
			return 0, "1. exit with 0"
		},
		func(name string, args ...string) (statusCode int, output string) {
			return 0, "2. exit with 0"
		},
		commandtest.CallTimesTestHelperProcess(
			func(name string, args ...string) (statusCode int, output string) {
				return 0, "3.1. exit with 0"
			},
			func(name string, args ...string) (statusCode int, output string) {
				return 1, "3.2. exit with 1"
			},
			func(name string, args ...string) (statusCode int, output string) {
				return 0, "3.3. exit with 0"
			},
		),
		commandtest.CallTimesTestHelperProcess(
			func(name string, args ...string) (statusCode int, output string) {
				return 0, "4.1. exit with 0"
			},
			func(name string, args ...string) (statusCode int, output string) {
				return 0, "4.2. exit with 0"
			},
			func(name string, args ...string) (statusCode int, output string) {
				return 2, "4.3. exit with 2"
			},
		),
	)
}
