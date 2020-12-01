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
	return func(options *testingCommanderOptions) {
		options.HelperName = helperName
	}
}

func WithCommanderEnvs(envs func(prevEnvs []string) []string) TestingCommanderOption {
	return func(options *testingCommanderOptions) {
		options.Envs = envs
	}
}

type testingCommanderOptions struct {
	HelperName string
	Envs       func(prevEnvs []string) []string
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
			envs = options.Envs(envs)
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
