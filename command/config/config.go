package config

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/repository"
)

// GetConfig get config
func GetConfig(key string, r repository.ConfigReader) (string, error) {
	if key == "" {
		return "", errors.NewError(
			errors.WithCode(exitcode.MissingArguments),
			errors.WithMsg("config key required"),
		)
	}
	v, err := r.GetConfig(key)
	if err != nil {
		return "", err
	}

	if v == "" {
		return "", errors.NewError(errors.WithCode(exitcode.RepoConfigNotFound), errors.WithMsgf("config not found, key:%s", key))
	}

	return v, nil
}

// UpdateConfig update config
func UpdateConfig(key, value string, r repository.ConfigWriter) error {
	if key == "" {
		return errors.NewError(
			errors.WithCode(exitcode.MissingArguments),
			errors.WithMsg("config key required"),
		)
	}

	return r.SetConfig(key, value)
}
