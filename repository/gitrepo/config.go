package gitrepo

import (
	"fmt"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/fs"
	"github.com/funnyecho/git-syncer/repository/gitrepo/gitter"
)

const (
	configKeyPrefix   = "git-syncer"
	projectConfigName = ".git-syncer-config"
)

func (r *repo) GetConfig(keys ...string) (string, error) {
	return r.getConfig(keys)
}

func (r *repo) SetConfig(key, value string) error {
	panic("implement me")
}

func (r *repo) getConfig(keys []string) (val string, err error) {
	pcExisted, _ := fs.IsFileExists(projectConfigName)
	remote := r.remote

	defer func() {
		if val != "" {
			err = nil
		}
	}()

	for _, key := range keys {
		if remote != "" {
			if pcExisted {
				val, err = r.gitter.ConfigGet(
					gitter.WithConfigGetFromFile(projectConfigName),
					gitter.WithConfigGetKey(fmt.Sprintf("%s.%s.%s", configKeyPrefix, remote, key)),
				)
				if val != "" {
					return
				}
			}

			val, err = r.gitter.ConfigGet(
				gitter.WithConfigGetKey(fmt.Sprintf("%s.%s.%s", configKeyPrefix, remote, key)),
			)
			if val != "" {
				return
			}
		}

		if pcExisted {
			val, err = r.gitter.ConfigGet(
				gitter.WithConfigGetFromFile(projectConfigName),
				gitter.WithConfigGetKey(fmt.Sprintf("%s.%s", configKeyPrefix, key)),
			)
			if val != "" {
				return
			}
		}

		val, err = r.gitter.ConfigGet(
			gitter.WithConfigGetKey(fmt.Sprintf("%s.%s", configKeyPrefix, key)),
		)
		if val != "" {
			return
		}
	}

	return "", errors.NewError(
		errors.WithStatusCode(exitcode.Usage),
		errors.WithMsg(fmt.Sprintf("failed to get config: %v", keys)),
	)
}
