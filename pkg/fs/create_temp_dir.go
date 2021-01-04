package fs

import (
	"fmt"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"io/ioutil"
)

type tempDirFallback = struct {
	dir     string
	pattern string
}

type TempDirFallbackOption = func() tempDirFallback
type tempDirFallbackHandler = func(dir, pattern string) (string, error)

func CreateTempDir(fallbacks ...TempDirFallbackOption) (string, []error) {
	return iterateTempDirFallbacks(func(dir, pattern string) (string, error) {
		return ioutil.TempDir(dir, pattern)
	}, fallbacks...)
}

func WithTempDirFallback(dir, pattern string) TempDirFallbackOption {
	return func() tempDirFallback {
		return tempDirFallback{
			dir,
			pattern,
		}
	}
}

func withDefaultTempDirFallback() TempDirFallbackOption {
	return func() tempDirFallback {
		return tempDirFallback{
			dir:     "",
			pattern: "*",
		}
	}
}

func iterateTempDirFallbacks(handler tempDirFallbackHandler, fallbacks ...TempDirFallbackOption) (string, []error) {
	if len(fallbacks) == 0 {
		fallbacks = append([]TempDirFallbackOption{}, withDefaultTempDirFallback())
	}

	var errStack []error

	for _, fallbackOption := range fallbacks {
		fallback := fallbackOption()
		dir, err := handler(fallback.dir, fallback.pattern)
		if err == nil && dir != "" {
			return dir, nil
		} else {
			errStack = append(errStack, errors.NewError(
				errors.WithMsg(fmt.Sprintf("fallback %v", fallbacks)),
				errors.WithErr(err),
			))
		}
	}

	return "", errStack
}
