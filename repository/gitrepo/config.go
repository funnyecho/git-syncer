package gitrepo

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
)

const (
	// ProjectConfigName project leve config without prefix
	ProjectConfigName = ".git-syncer-config"
)

const (
	// ConfigSyncRoot sync_root key in config file
	ConfigSyncRoot = "syncRoot"
)

func (r *repo) GetConfig(key string) (string, error) {
	return r.getConfig(key)
}

func (r *repo) SetConfig(key, value string) error {
	return r.setConfig(key, value)
}

func (r *repo) getConfig(key string) (val string, err error) {
	remote := r.remote

	defer func() {
		if val != "" {
			err = nil
		}
	}()

	if remote != "" {
		val, err = r.gitter.ConfigGet(
			fmt.Sprintf("%s.%s", remote, key),
			gitter.ConfigGetOptions{
				File: ProjectConfigName,
			},
		)
		if val != "" {
			return
		}
	}

	val, err = r.gitter.ConfigGet(
		key,
		gitter.ConfigGetOptions{
			File: ProjectConfigName,
		},
	)
	if val != "" {
		return
	}

	val, err = r.gitter.ConfigGet(key, gitter.ConfigGetOptions{})
	if val != "" {
		return
	}

	return "", errors.NewError(
		errors.WithCode(exitcode.Usage),
		errors.WithMsg(fmt.Sprintf("failed to get config: %s", key)),
	)
}

func (r *repo) setConfig(key, value string) error {
	remote := r.remote

	if remote != "" {
		return r.gitter.ConfigSet(
			fmt.Sprintf("%s.%s", remote, key),
			value,
			gitter.ConfigSetOptions{
				File: ProjectConfigName,
			},
		)
	}

	return r.gitter.ConfigSet(
		key,
		value,
		gitter.ConfigSetOptions{
			File: ProjectConfigName,
		},
	)
}
