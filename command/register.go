package command

import (
	"github.com/funnyecho/git-syncer/command/catchup"
	"github.com/funnyecho/git-syncer/command/config"
	"github.com/funnyecho/git-syncer/command/push"
	"github.com/funnyecho/git-syncer/command/setup"
	"github.com/mitchellh/cli"
)

// Register all syner command factories
func Register() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"setup":   setup.Factory,
		"catchup": catchup.Factory,
		"push":    push.Factory,
		"config":  config.Factory,
		//"add-scope":    add_scope.Factory,
		//"remove-scope": remove_scope.Factory,
		//"show":         show.Factory,
		//"log":          log.Factory,
	}
}
