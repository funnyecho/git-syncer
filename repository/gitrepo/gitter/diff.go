package gitter

import (
	"bufio"
	"bytes"
	"fmt"
)

// Diff from specified commit
type Diff interface {
	DiffAM(path string, commit string) ([]string, error)
	DiffD(path string, commit string) ([]string, error)
}

func (g *git) DiffAM(path string, commit string) ([]string, error) {
	return g.diff(path, commit, diffOptions{"AM"})
}

func (g *git) DiffD(path string, commit string) ([]string, error) {
	return g.diff(path, commit, diffOptions{"D"})
}

type diffOptions struct {
	filter string
}

func (g *git) diff(path string, commit string, options diffOptions) ([]string, error) {
	args := []string{"diff", "--name-only", "--no-renames"}

	if options.filter != "" {
		args = append(args, fmt.Sprintf("--diff-filter=%s", options.filter))
	}

	if path == "" {
		path = "."
	}

	args = append(args, commit, "--", path)

	var files []string

	cmd := g.command("git", args...)

	v, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(v))
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	return files, err
}
