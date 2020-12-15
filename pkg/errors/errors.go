package errors

import (
	"fmt"
)

// [Working with Errors in Go 1.13](https://blog.golang.org/go1.13-errors)
type Error struct {
	StatusCode int
	Err error
	Msg string
}

func (e *Error) Unwrap() error {
	return e.Err
}

// No need to print wrapped error, `go` print it automatically
func (e *Error) Error() string {
	if e.Msg == "" {
		return fmt.Sprintf("status code: %d", e.StatusCode)
	} else {
		return fmt.Sprintf("status code: %d, error msg: %s", e.StatusCode, e.Msg)
	}
}

func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}

	return t.StatusCode >= 0 && e.StatusCode == t.StatusCode
}

func (e *Error) As(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}

	return t.StatusCode >= 0
}

type errorOption = func(err *Error)

func NewError(options ...errorOption) *Error {
	err := &Error{}
	for _, option := range options {
		option(err)
	}

	return err
}

func WithStatusCode(code int) errorOption {
	return func(err *Error) {
		err.StatusCode = code
	}
}

func WithMsg(msg string) errorOption {
	return func(err *Error) {
		err.Msg = msg
	}
}

func WithErr(wrapErr error) errorOption {
	return func(err *Error) {
		err.Err = wrapErr
	}
}