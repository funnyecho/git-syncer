package git

import (
	"strings"
)

// ProjectConfigPath project level config file
const ProjectConfigPath = "./.git-syncer-config"

// GetConfig implement GetConfig
func (g *Git) GetConfig(key string) (string, error) {
	args := []string{"config", "-f", ProjectConfigPath}

	args = append(args, key)

	v, err := output(args)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(v)), nil
}

// SetConfig implement SetConfig
func (g *Git) SetConfig(key, value string) error {
	args := []string{"config", "-f", ProjectConfigPath}

	args = append(args, key, value)

	_, err := output(args)
	return err
}
