package command

import "os/exec"

type Commander = func(name string, args ...string) *exec.Cmd

var command = exec.Command

func UseCommand() Commander {
	return command
}

func PopCommand(cmd Commander) {
	command = cmd
}

func PushCommand(cmd Commander) Commander {
	defer func() {
		command = cmd
	}()

	return command
}
