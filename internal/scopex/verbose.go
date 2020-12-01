package scopex

import (
	"context"
	"github.com/funnyecho/git-syncer/internal/constants"
)

const (
	VerboseSilent = iota
	VerboseError
	VerboseInfo
	VerboseDebug
)

const LowestVerbose = VerboseDebug

func UseVerbose(ctx context.Context) int {
	val := ctx.Value(constants.ContextKeyVerbose)
	if val == nil {
		return VerboseSilent
	} else {
		return min(LowestVerbose, val.(int))
	}
}

func WithVerbose(ctx context.Context, verbose int) context.Context {
	return context.WithValue(ctx, constants.ContextKeyVerbose, min(LowestVerbose, verbose))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}