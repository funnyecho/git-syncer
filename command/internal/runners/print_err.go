package runners

import (
	"fmt"
	"os"

	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/mitchellh/cli"
)

// PrintErr print error
func PrintErr(_ []string) (runner.BubbleTap, error) {
	return func(e error) error {
		if e != nil {
			ui := cli.ColoredUi{
				ErrorColor: cli.UiColorRed,
				Ui: &cli.BasicUi{
					ErrorWriter: os.Stderr,
				},
			}
			ui.Error(fmt.Sprint(e))
		}
		return e
	}, nil
}
