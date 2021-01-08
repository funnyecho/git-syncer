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

func WithConfigGetKey(key string) WithArgs {
	return func() string {
		return fmt.Sprintf("--get \"%s\"", key)
	}
}

func WithConfigGetFromFile(filePath string) WithArgs {
	return func() string {
		return fmt.Sprintf("-f '%s'", filePath)
	}
}
