package errors

import (
	"errors"
	"fmt"

	"github.com/funnyecho/git-syncer/constants/exitcode"
)

// Error stacked error with status code
// [Working with Errors in Go 1.13](https://blog.golang.org/go1.13-errors)
type Error struct {
	Code int
	Err  error
	Msg  string
}

// Unwrap implement `Unwrap`
func (e *Error) Unwrap() error {
	return e.Err
}

// No need to print wrapped error, `go` print it automatically
func (e *Error) Error() string {
	if e.Msg == "" {
		return fmt.Sprintf("status code: %d", e.Code)
	}

	return fmt.Sprintf("status code: %d, error msg: %s", e.Code, e.Msg)
}

// Is implement `Is`
func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}

	eCode := GetErrorCode(e)
	tCode := GetErrorCode(t)

	return tCode > 0 && eCode == tCode
}

// As implement `As`
func (e *Error) As(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}

	return GetErrorCode(t) > 0
}

type errorOption = func(err *Error)

// NewError create error
func NewError(options ...errorOption) *Error {
	err := &Error{}
	for _, option := range options {
		option(err)
	}

	return err
}

// WithCode with status code
func WithCode(code int) errorOption {
	return func(err *Error) {
		err.Code = code
	}
}

// WithMsg with error msg
func WithMsg(msg string) errorOption {
	return func(err *Error) {
		err.Msg = msg
	}
}

// WithErr with wrap error
func WithErr(wrapErr error) errorOption {
	return func(err *Error) {
		err.Err = wrapErr
	}
}

// GetErrorCode get status code of error
func GetErrorCode(err error) int {
	if err == nil {
		return exitcode.Unknown
	}

	var errWithCode *Error

	hasCode := errors.As(err, &errWithCode)
	if !hasCode {
		return exitcode.Unknown
	}

	code := errWithCode.Code
	if code > 0 {
		return code
	}

	return GetErrorCode(errWithCode.Unwrap())
}
