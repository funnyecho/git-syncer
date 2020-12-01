package scopex_test

import (
	"context"
	"github.com/funnyecho/git-syncer/internal/scopex"
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestProjectDir(t *testing.T) {
	ctx := context.Background()

	Empty(t, scopex.UseProjectDir(ctx))

	mockDir := "foo/bar"
	ctx = scopex.WithProjectDir(ctx, mockDir)

	Equal(t, mockDir, scopex.UseProjectDir(ctx))
}