package syncer

import (
	"fmt"
	"github.com/funnyecho/git-syncer/pkg/fs"
	"github.com/funnyecho/git-syncer/pkg/gitter"
)

const (
	configKeyPrefix = "git-syncer"
	projectConfigName = ".git-syncer-config"
)

func GetConfig(key string) (string, error) {
	return getConfig(key)
}

func MustGetConfig(key string) string {
	v, err := getConfig(key)

	if err != nil {
		panic(err)
	}

	return v
}

func getConfig(key string) (val string, err error) {
	pcExisted, _ := fs.IsFileExists(projectConfigName)

	if remote != "" {
		if pcExisted {
			val, err = gitter.ConfigGet(
				gitter.WithConfigGetFromFile(projectConfigName),
				gitter.WithConfigGetKey(fmt.Sprintf("%s.%s.%s", configKeyPrefix, remote, key)),
			)
			if err != nil || val != "" {
				return
			}
		}

		val, err = gitter.ConfigGet(
			gitter.WithConfigGetKey(fmt.Sprintf("%s.%s.%s", configKeyPrefix, remote, key)),
		)
		if err != nil || val != "" {
			return
		}
	}

	if pcExisted {
		val, err = gitter.ConfigGet(
			gitter.WithConfigGetFromFile(projectConfigName),
			gitter.WithConfigGetKey(fmt.Sprintf("%s.%s", configKeyPrefix, key)),
		)
		if err != nil || val != "" {
			return
		}
	}

	val, err = gitter.ConfigGet(
		gitter.WithConfigGetKey(fmt.Sprintf("%s.%s", configKeyPrefix, key)),
	)

	return
}
