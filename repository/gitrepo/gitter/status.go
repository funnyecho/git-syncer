package gitter

import (
	"bufio"

	"github.com/funnyecho/git-syncer/pkg/command"
)

// GetPorcelainStatus returning the short format status of repo
func (g *git) GetPorcelainStatus(withArgs ...WithArgs) (status []string, err error) {
	return g.getPorcelainStatus(withArgs...)
}

// WithUnoPorcelainStatus specify `-uno` argument
func (g *git) WithUnoPorcelainStatus() WithArgs {
	return func() string {
		return "-uno"
	}
}

func (g *git) getPorcelainStatus(withArgs ...WithArgs) (status []string, err error) {
	args := []string{"status", "--porcelain"}

	for _, fn := range withArgs {
		arg := fn()
		args = append(args, arg)
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

	return status, nil
}
