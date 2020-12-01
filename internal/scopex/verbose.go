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

func UseVerbose(ctx context.Context) int {
	return  ctx.Value(constants.ContextKeyVerbose).(int)
}

func WithVerbose(ctx context.Context, verbose int) context.Context {
	return context.WithValue(ctx, constants.ContextKeyVerbose, verbose)
}
