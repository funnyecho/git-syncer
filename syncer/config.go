package syncer

import (
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer/gitter"
)

// Config command handler
func Config(git gitter.ConfigReadWriter, args []string) error {
	if git == nil {
		return errors.Err(exitcode.InvalidParams, "missing `Gitter`")
	}

	if args == nil {
		return errors.Err(exitcode.InvalidParams, "missing arguments")
	}

	if argsL := len(args); argsL == 0 || argsL > 2 {
		return errors.Err(exitcode.Usage, "invalid config args length")
	}

	key := args[0]
	if len(args) == 1 {
		value, err := git.GetConfig(key)
		if err != nil {
			return errors.Wrap(err, "config not found")
		}

		if value == "" {
			return errors.Err(exitcode.RepoConfigNotFound, "config not found")
		}

		fmt.Println(value)
		return nil
	}

	value := args[1]
	if err := git.SetConfig(key, value); err != nil {
		return errors.Wrap(err, "config update failed")
	}

	return nil
}

func getConfig(configRW gitter.ConfigReadWriter, key string) (string, error) {
	return configRW.GetConfig(key)
}

func getSyncRootConfig(configRW gitter.ConfigReadWriter) string {
	v, err := getConfig(configRW, "sync-root")
	if err != nil || v == "" {
		v = "."
	}

	return v
}
