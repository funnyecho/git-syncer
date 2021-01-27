package runners

import (
	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/pkg/log"
)

// WorkingHead checkout to new working head if head flag provided
func WorkingHead(_ []string) (runner.BubbleTap, error) {
	head, _ := useFlag(flagWorkingHead)
	if head == "" {
		return nil, nil
	}

	git, gitErr := UseGitter()
	if gitErr != nil {
		return nil, gitErr
	}

	oriHead, pushErr := git.PushHead(head)
	if pushErr != nil {
		return nil, pushErr
	}

	return func(e error) error {
		popErr := git.PopHead(oriHead)
		if popErr != nil {
			log.Errore("failed to reset to head", popErr, "head", oriHead)
		}

		return e
	}, nil
}
