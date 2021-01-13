package gitter

import (
	"fmt"
)

// ConfigGetOptions config getter options
type ConfigGetOptions struct {
	File string
}

// ConfigSetOptions config setter options
type ConfigSetOptions struct {
	File string
}

// Config config reader and writer
type Config interface {
	ConfigGet(key string, options ConfigGetOptions) (string, error)
	ConfigSet(key, value string, options ConfigSetOptions) error
}

func (g *git) ConfigGet(key string, options ConfigGetOptions) (string, error) {
	args := []string{"config"}

	if options.File != "" {
		args = append(args, fmt.Sprintf("-f %s", options.File))
	}

	args = append(args, fmt.Sprintf("--get %s", key))

	cmd := g.command("git", args...)

	v, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(v), nil
}

func (g *git) ConfigSet(key, value string, options ConfigSetOptions) error {
	args := []string{"config"}

	if options.File != "" {
		args = append(args, fmt.Sprintf("-f %s", options.File))
	}

	args = append(args, fmt.Sprintf("%s %s", key, value))

	cmd := g.command("git", args...)

	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
