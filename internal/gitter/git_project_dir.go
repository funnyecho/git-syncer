package gitter

import (
	"bytes"
	"os/exec"
)

func GetProjectDir() (dir string, err error)  {
	cmd := exec.Command("git", "rev-parse --show-toplevel")

	var stdout bytes.Buffer

	cmd.Stdout= &stdout

	err = cmd.Run()
	if err != nil {
		return
	}

	dir = string(stdout.Bytes())
	return
}
