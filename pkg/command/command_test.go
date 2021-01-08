package command_test

import (
	"github.com/funnyecho/git-syncer/pkg/command"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"reflect"
	"runtime"
	"testing"
)

func TestCommandStack(t *testing.T) {
	assert.NotNil(t, command.UseCommand())

	fakeCommand := func(name string, args ...string) *exec.Cmd {
		return exec.Command(name, args...)
	}

	oriCommand := command.PushCommand(fakeCommand)
	compareCommander(t, fakeCommand, command.UseCommand())

	command.PopCommand(oriCommand)
	compareCommander(t, oriCommand, command.UseCommand())
}

// refer to: [the work around for comparing function "equality"](https://github.com/stretchr/testify/issues/182#issuecomment-495359313)
func compareCommander(t *testing.T, func1 command.Commander, func2 command.Commander) {
	funcName1 := runtime.FuncForPC(reflect.ValueOf(func1).Pointer()).Name()
	funcName2 := runtime.FuncForPC(reflect.ValueOf(func2).Pointer()).Name()
	assert.Equal(t, funcName1, funcName2)
}
