package variable

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
)

var (
	// ErrNotFound variable not found error
	ErrNotFound = errors.Err(exitcode.VariableNotFound, "variable not found")
)
