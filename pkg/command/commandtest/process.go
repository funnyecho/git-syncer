package commandtest

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type CommandExecutor = func(name string, args ...string) (statusCode int, output string)

func TestHelperProcess(t *testing.T, executors ...CommandExecutor) {
	if !EnvUseGoWantHelperProcess() {
		return
	}

	defer os.Exit(0)

	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	cmd, args := args[0], args[1:]

	var code int
	var stdOuts []string
	var stdErr string
	var output string

	for _, executor := range executors {
		if executor == nil {
			continue
		}

		code, output = executor(cmd, args...)
		if code != 0 {
			stdErr = output
			break
		} else {
			stdOuts = append(stdOuts, output)
		}
	}

	fmt.Fprint(os.Stdout, strings.Join(stdOuts, "\n"))
	fmt.Fprint(os.Stderr, stdErr)
	os.Exit(code)
}

func CallTimesTestHelperProcess(executors ...CommandExecutor) CommandExecutor {
	return func(name string, args ...string) (statusCode int, output string) {
		if executors == nil || len(executors) == 0 {
			return
		}

		maxExecCursor := len(executors)

		execTimes := EnvUseCommandTimes()
		execCursor := execTimes
		if execCursor > maxExecCursor {
			execCursor = maxExecCursor
		}

		return executors[execCursor-1](name, args...)
	}
}
