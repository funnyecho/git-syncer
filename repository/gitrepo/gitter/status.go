package gitter

import (
	"bufio"
)

// statusOptions git status options
type statusOptions struct {
	Uno bool
}

// Status git status reader interface
type Status interface {
	GetPorcelainStatus() (status []string, err error)
	GetUnoPorcelainStatus() (status []string, err error)
}

// GetPorcelainStatus returning the short format status of repo
func (g *git) GetPorcelainStatus() (status []string, err error) {
	return g.getPorcelainStatus(statusOptions{})
}

// GetUnoPorcelainStatus returning the short format status of repo with `-uno` argument
func (g *git) GetUnoPorcelainStatus() (status []string, err error) {
	return g.getPorcelainStatus(statusOptions{
		Uno: true,
	})
}

func (g *git) getPorcelainStatus(options statusOptions) (status []string, err error) {
	args := []string{"status", "--porcelain"}

	if options.Uno {
		args = append(args, "-uno")
	}

	cmd := g.command("git", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			status = nil
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			status = append(status, scanner.Text())
		}
	}()

	err = cmd.Wait()

	return status, err
}
