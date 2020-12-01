package scopex_test

import (
	"context"
	"github.com/funnyecho/git-syncer/internal/scopex"
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestLowestVerbose(t *testing.T) {
	Equal(
		t,
		scopex.VerboseDebug,
		scopex.LowestVerbose,
		"when change verbose level, correct the lowest verbose level",
	)
}

func TestVerbose(t *testing.T) {
	ctx := context.Background()
	Equal(t, scopex.VerboseSilent, scopex.UseVerbose(ctx))

	ctx = scopex.WithVerbose(ctx, scopex.VerboseDebug)
	Equal(t, scopex.VerboseDebug, scopex.UseVerbose(ctx))

	ctx = scopex.WithVerbose(ctx, scopex.LowestVerbose + 1)
	Equal(t, scopex.LowestVerbose, scopex.UseVerbose(ctx), "verbose is always smaller than lowest verbose level")
}
