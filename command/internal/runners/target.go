package runners

import "github.com/funnyecho/git-syncer/command/internal/runner"

// WithTarget runner without bubble tap
func WithTarget(f func(args []string) error) runner.CaptureTap {
	return func(args []string) (runner.BubbleTap, error) {
		return nil, f(args)
	}
}
