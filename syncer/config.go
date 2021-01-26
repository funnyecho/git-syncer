package syncer

import (
	"flag"
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/syncer/gitter"
)

// ConfigParams handler params
type ConfigParams struct {
	gitter.ConfigReadWriter
	*flag.FlagSet
}

// Config command handler
func Config(params ConfigParams) error {
	if params.ConfigReadWriter == nil {
		return errors.Err(exitcode.InvalidParams, "missing `Gitter`")
	}

	args := params.Args()

	if args == nil {
		return errors.Err(exitcode.MissingArguments, "missing arguments")
	}

	if len(args) > 2 {
		return errors.Err(exitcode.Usage, "invalid config args length")
	}

	key := args[0]
	if len(args) == 1 {
		value, err := params.GetConfig(key)
		if err != nil {
			return errors.WrapC(err, exitcode.RepoConfigNotFound, "config not found")
		}

		if value == "" {
			return errors.Err(exitcode.RepoConfigNotFound, "config not found")
		}

		fmt.Println(value)
		return nil
	}

	value := args[1]
	if err := params.SetConfig(key, value); err != nil {
		return errors.WrapC(err, exitcode.RepoUnknown, "config update failed")
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
