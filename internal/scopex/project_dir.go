package scopex

import (
	"context"
	"github.com/funnyecho/git-syncer/internal/constants"
)

func UseProjectDir(ctx context.Context) string {
	val := ctx.Value(constants.ContextKeyProjectDir)

	if val == nil {
		return ""
	} else {
		return val.(string)
	}
}

func WithProjectDir(ctx context.Context, dir string) context.Context  {
	return context.WithValue(ctx, constants.ContextKeyProjectDir, dir)
}
