package gitter

import (
	"bufio"
	"bytes"
)

func (g *git) ListFiles(path string) ([]string, error) {
	if path == "" {
		path = "."
	}
	cmd := g.command("git", "ls-files", "--", path)

	v, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var fs []string

	scanner := bufio.NewScanner(bytes.NewBuffer(v))
	for scanner.Scan() {
		fs = append(fs, scanner.Text())
	}

	return fs, nil
}
