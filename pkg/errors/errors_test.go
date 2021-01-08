package errors_test

import (
	stdErr "errors"
	"fmt"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestErrors(t *testing.T) {
	plainErr := fmt.Errorf("plain err")
	plainErr2 := fmt.Errorf("plain err 2")

	unknownErrWithWrappedErr := &errors.Error{
		StatusCode: exitcode.Unknown,
		Err:        plainErr,
		Msg:        "unknown err with wrapped error",
	}

	unknownErrWithoutWrappedErr := &errors.Error{
		StatusCode: exitcode.Unknown,
		Msg:        "unknown err without wrapped error",
	}

	usageErrWithWrappedErr := &errors.Error{
		StatusCode: exitcode.Usage,
		Err:        plainErr,
		Msg:        "usage err with wrapped error",
	}

	plainErrWithWrappedError := fmt.Errorf("plain err 3 with usage error: %w", usageErrWithWrappedErr)

	t.Run("errors.Is", func(t *testing.T) {
		True(t, stdErr.Is(unknownErrWithWrappedErr, unknownErrWithoutWrappedErr), "status code equal")
		False(t, stdErr.Is(unknownErrWithWrappedErr, usageErrWithWrappedErr), "status code not equal")

		True(t, stdErr.Is(usageErrWithWrappedErr, plainErr), "with wrapped plain err")
		False(t, stdErr.Is(usageErrWithWrappedErr, plainErr2), "wrapped plain err not equal")

		True(t, stdErr.Is(plainErrWithWrappedError, usageErrWithWrappedErr), "plain err wrapped error")
		False(t, stdErr.Is(plainErrWithWrappedError, unknownErrWithWrappedErr), "wrapped plain error not equal")
	})

	t.Run("errors.As", func(t *testing.T) {
		var sErr error
		var iErr *errors.Error

		True(t, stdErr.As(unknownErrWithoutWrappedErr, &sErr), "implemented standard error interface")
		NotNil(t, sErr)
		NotNil(t, sErr.(*errors.Error))

		False(t, stdErr.As(plainErr, &iErr))
		Nil(t, iErr)

		True(t, stdErr.As(plainErrWithWrappedError, &iErr))
		NotNil(t, iErr)
		Equal(t, exitcode.Usage, iErr.StatusCode)
	})

	t.Run("errors.Unwrap", func(t *testing.T) {
		Nil(t, stdErr.Unwrap(unknownErrWithoutWrappedErr))
		NotNil(t, stdErr.Unwrap(unknownErrWithWrappedErr))
	})
}

func TestErrorsConstructor(t *testing.T) {
	plainErr := fmt.Errorf("plain err")
	err := errors.NewError()

	Equal(t, exitcode.Nil, err.StatusCode)
	Empty(t, err.Msg)
	Nil(t, err.Err)
	Nil(t, err.Unwrap())

	err = errors.NewError(
		errors.WithErr(plainErr),
		errors.WithStatusCode(exitcode.MissingArguments),
		errors.WithMsg("foobar"),
	)

	Equal(t, exitcode.MissingArguments, err.StatusCode)
	Equal(t, "foobar", err.Msg)
	Equal(t, plainErr, err.Err)
	Equal(t, plainErr, err.Unwrap())
}
