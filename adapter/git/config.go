package git

import "fmt"

// ProjectConfigPath project level config file
const ProjectConfigPath = "./.git-syncer-config"

// GitConfigPrefix prefix in git config
const GitConfigPrefix = "git-syncer"

// GetConfig implement GetConfig
func (g *Git) GetConfig(key string) (string, error) {
	args := []string{"config", "-f", ProjectConfigPath}

	args = append(args, key)

	v, err := output(args)
	if err == nil {
		return v, nil
	}

	v, err = output([]string{"config", fmt.Sprintf("%s.%s", GitConfigPrefix, key)})
	if err == nil {
		return v, nil
	}

	v, err = output([]string{"config", key})
	if err == nil {
		return v, nil
	}

	return "", err
}

// SetConfig implement SetConfig
func (g *Git) SetConfig(key, value string) error {
	args := []string{"config", "-f", ProjectConfigPath}

	args = append(args, key, value)

	_, err := output(args)
	return err
}
