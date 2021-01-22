package command

import (
	"flag"
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/repository"
	"github.com/mitchellh/cli"
)

// ConfigFactory of command `push`
func ConfigFactory() (cli.Command, error) {
	return &configCmd{}, nil
}

type configCmd struct {
}

func (c *configCmd) Help() string {
	return "Config getter and setter"
}

func (c *configCmd) Synopsis() string {
	return "Config getter and setter"
}

func (c *configCmd) Run(args []string) int {
	var opt BasicOptions
	fs := flag.NewFlagSet("config", flag.ContinueOnError)

	if optCode := OptionsBinder(&opt, fs)(args); optCode != exitcode.Nil {
		return optCode
	}

	repo, repoCode := NewRepo(opt.Base, opt.Remote)
	if repoCode != exitcode.Nil {
		return repoCode
	}

	v := fs.Args()
	if len(v) == 0 {
		log.Errorw("failed to get key or value")
		return exitcode.Usage
	}

	if len(v) == 1 {
		value, configErr := ExecGetConfig(v[0], repo)
		if configErr != nil {
			log.Errore("failed to get config", configErr, "key", v[0])
			return errors.GetErrorCode(configErr)
		}

		fmt.Print(value)
		return exitcode.Nil
	}

	configErr := ExecUpdateConfig(v[0], v[1], repo)
	if configErr != nil {
		fmt.Println("failed to update config", configErr, "key", v[0], "value", v[1])
		return errors.GetErrorCode(configErr)
	}

	return exitcode.Nil
}

// ExecGetConfig get config
func ExecGetConfig(key string, r repository.ConfigReader) (string, error) {
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

// ExecUpdateConfig update config
func ExecUpdateConfig(key, value string, r repository.ConfigWriter) error {
	if key == "" {
		return errors.NewError(
			errors.WithCode(exitcode.MissingArguments),
			errors.WithMsg("config key required"),
		)
	}

	return r.SetConfig(key, value)
}
