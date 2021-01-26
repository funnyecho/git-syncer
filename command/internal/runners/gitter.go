package runners

import (
	"fmt"

	"github.com/funnyecho/git-syncer/adapter/git"
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/syncer/gitter"
)

var gt gitter.Gitter

var defaultRemote = "default"

// Gitter tap to gitter
func Gitter(_ []string) (runner.BubbleTap, error) {
	remote, remoteErr := UseFlagRemote()

	if remoteErr != nil {
		return nil, remoteErr
	}

	gt = &remoteGitter{
		remote,
		git.New(),
	}

	return nil, nil
}

// UseGitter get gitter
func UseGitter() (gitter.Gitter, error) {
	if gt == nil {
		return nil, errors.Err(exitcode.InvalidRunnerDependency, "failed to get gitter from runner")
	}

	return gt, nil
}

type remoteGitter struct {
	remote string
	gitter.Gitter
}

func (g *remoteGitter) GetConfig(key string) (string, error) {
	if g.remote != "" {
		v, err := g.Gitter.GetConfig(fmt.Sprintf("%s.%s", g.remote, key))
		if err == nil {
			return v, nil
		}
		log.Infow("failed to get config in remote", "remote", g.remote, "key", key)
	}

	v, err := g.Gitter.GetConfig(fmt.Sprintf("%s.%s", defaultRemote, key))
	if err == nil {
		return v, nil
	}
	log.Infow("failed to get config in remote", "remote", defaultRemote, "key", key)

	v, err = g.Gitter.GetConfig(key)
	if err == nil {
		return v, nil
	}
	log.Infow("failed to get config without remote", "key", key)

	return "", errors.Err(exitcode.RepoConfigNotFound, "failed to get config %s", key)
}

// SetConfig implement SetConfig
func (g *remoteGitter) SetConfig(key, value string) error {
	if g.remote != "" {
		err := g.Gitter.SetConfig(fmt.Sprintf("%s.%s", g.remote, key), value)
		if err == nil {
			return err
		}
	}

	return g.Gitter.SetConfig(fmt.Sprintf("%s.%s", defaultRemote, key), value)
}
