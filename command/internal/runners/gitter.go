package runners

import (
	"fmt"

	"github.com/funnyecho/git-syncer/adapter/git"
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
	"github.com/funnyecho/git-syncer/syncer/gitter"
	"github.com/funnyecho/git-syncer/variable"
)

var defaultRemote = "default"

// Gitter tap to gitter
func Gitter(_ []string) (runner.BubbleTap, error) {
	remote, remoteErr := UseFlagRemote()

	if remoteErr != nil {
		return nil, remoteErr
	}

	variable.WithGitter(&remoteGitter{
		remote,
		git.New(),
	})

	return nil, nil
}

// UseGitter get gitter
var UseGitter = variable.UseGitter

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
		log.Debugw("config not found with remote", "remote", g.remote, "key", key)
	}

	v, err := g.Gitter.GetConfig(fmt.Sprintf("%s.%s", defaultRemote, key))
	if err == nil {
		return v, nil
	}
	log.Debugw("config not found with remote", "remote", defaultRemote, "key", key)

	v, err = g.Gitter.GetConfig(key)
	if err == nil {
		return v, nil
	}
	log.Debugw("config not found without remote", "key", key)

	return "", errors.Err(exitcode.RepoConfigNotFound, "failed to get config %s", key)
}

// SetConfig implement SetConfig
func (g *remoteGitter) SetConfig(key, value string) error {
	r := g.remote
	if r == "" {
		r = defaultRemote
	}

	log.Debugw("setting config into remote", "remote", r, "key", key, "value", value)
	return g.Gitter.SetConfig(fmt.Sprintf("%s.%s", r, key), value)
}
