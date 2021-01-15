package gitter

import (
	"bufio"
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

	args = append(args, commit, fmt.Sprintf("-- \"%s\"", path))

	var files []string

	cmd := g.command("git", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			files = nil
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			files = append(files, scanner.Text())
		}
	}()

	err = cmd.Wait()

	return files, err
}
