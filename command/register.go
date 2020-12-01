package command

import (
	"github.com/funnyecho/git-syncer/command/setup"
	"github.com/mitchellh/cli"
)

func Register() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"setup":         setup.Factory,
		//"catchup":      catchup.Factory,
		//"push":         push.Factory,
		//"add-scope":    add_scope.Factory,
		//"remove-scope": remove_scope.Factory,
		//"show":         show.Factory,
		//"log":          log.Factory,
	}
}
