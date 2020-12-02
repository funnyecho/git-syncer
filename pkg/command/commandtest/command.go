package commandtest

import (
	"fmt"
	"github.com/funnyecho/git-syncer/pkg/command"
	"os"
	"os/exec"
	"strings"
)

type TestingCommanderOption = func(options *testingCommanderOptions)

func WithCommanderHelperName(helperName string) TestingCommanderOption {
	if !strings.HasPrefix(helperName, "Test") {
		panic("helper name must start with `Test`")
	}

	return func(options *testingCommanderOptions) {
		options.HelperName = helperName
	}
}

func WithCommanderEnvs(envs EnvFactory) TestingCommanderOption {
	return func(options *testingCommanderOptions) {
		options.Envs = envs
	}
}

func NewTestingCommander(withOptions ...TestingCommanderOption) command.Commander {
	options := &testingCommanderOptions{}
	for _, withOption := range withOptions {
		withOption(options)
	}

	helperName := options.HelperName
	if helperName == "" {
		helperName = "TestHelperProcess"
	}

	if !strings.HasPrefix(helperName, "Test") {
		panic("helper name must start with `Test`")
	}

	commandTimes := 0
	var envs []string

	return func(name string, args ...string) *exec.Cmd {
		commandTimes += 1
		if options.Envs != nil {
			envs = options.Envs(envs, name, args...)
		}

		cs := []string{fmt.Sprintf("-test.run=%s", helperName), "--", name}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = append(
			envs,
			EnvWithGoWantHelperProcess(),
			EnvWithCommandTimes(commandTimes),
		)

		return cmd
	}
}

type EnvFactory = func(prevEnvs []string, name string, args ...string) []string

type testingCommanderOptions struct {
	HelperName string
	Envs       EnvFactory
}
