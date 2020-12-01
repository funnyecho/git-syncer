package command

import "os/exec"

type Commander = func(name string, args ...string) *exec.Cmd

var Command = exec.Command

func PopCommand() Commander {
	return Command
}

func PushCommand(command Commander) {
	Command = command
}
