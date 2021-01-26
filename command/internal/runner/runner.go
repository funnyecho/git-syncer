package runner

import (
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/funnyecho/git-syncer/pkg/log"
)

// BubbleTap tap to bubble stage
type BubbleTap = func(error) error

// CaptureTap tap to capture stage
type CaptureTap = func(args []string) (BubbleTap, error)

// Run run tapper with args
func Run(args []string, tappers ...CaptureTap) (code int) {
	var bubbleTaps []BubbleTap
	var err error

	defer func() {
		pe := recover()
		if pe != nil {
			log.Errorw("uncapture error", "panic", pe)
			code = errors.GetErrorCode(pe.(error))
		}
	}()

	for _, tapper := range tappers {
		if tapper == nil {
			continue
		}

		bubbleTap, tapErr := tapper(args)
		if tapErr != nil {
			err = tapErr
			break
		}

		if bubbleTap != nil {
			bubbleTaps = append(bubbleTaps, bubbleTap)
		}
	}

	for i := len(bubbleTaps) - 1; i >= 0; i-- {
		err = bubbleTaps[i](err)
	}

	if err != nil {
		return errors.GetErrorCode(err)
	}

	return exitcode.Nil
}
