package gitter

import (
	"fmt"
)

func (g *git) ConfigGet(withArgs ...WithArgs) (string, error) {
	args := []string{"config"}

	for _, fn := range withArgs {
		arg := fn()
		args = append(args, arg)
	}

	cmd := g.command("git", args...)

	v, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(v), nil
}

func (g *git) ConfigSet(withArgs ...WithArgs) error {
	args := []string{"config"}

	for _, fn := range withArgs {
		arg := fn()
		args = append(args, arg)
	}

	cmd := g.command("git", args...)

	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

// WithConfigFile use config file
func WithConfigFile(filePath string) WithArgs {
	return func() string {
		return fmt.Sprintf("-f '%s'", filePath)
	}
}

// WithConfigSetKeyValue wrap to setter args
func WithConfigSetKeyValue(key, value string) WithArgs {
	return func() string {
		return fmt.Sprintf("\"%s\" %s", key, value)
	}
}

// WithConfigGetKey wrap to get getter args
func WithConfigGetKey(key string) WithArgs {
	return func() string {
		return fmt.Sprintf("--get \"%s\"", key)
	}
}
