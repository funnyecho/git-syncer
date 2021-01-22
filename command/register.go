package command

import (
	"github.com/mitchellh/cli"
)

// Register all syner command factories
func Register() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"setup":   SetupFactory,
		"catchup": CatchupFactory,
		"push":    PushFactory,
		"config":  ConfigFactory,
		//"add-scope":    add_scope.Factory,
		//"remove-scope": remove_scope.Factory,
		//"show":         show.Factory,
		//"log":          log.Factory,
	}
}
