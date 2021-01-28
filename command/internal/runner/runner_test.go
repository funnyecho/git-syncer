package runner_test

import (
	"testing"

	"github.com/funnyecho/git-syncer/command/internal/runner"
	"github.com/funnyecho/git-syncer/constants/exitcode"
	"github.com/funnyecho/git-syncer/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestRunner(t *testing.T) {
	t.Run("runners run in onion-like pattern", func(t *testing.T) {
		var output []int
		arguments := []string{"foo", "bar"}

		tap := func(i int) runner.CaptureTap {
			return func(args []string) (runner.BubbleTap, error) {

				assert.Equal(t, arguments, args)

				output = append(output, i)
				return func(e error) error {
					output = append(output, i*2)
					return e
				}, nil
			}
		}

		code := runner.Run(
			arguments,
			tap(1),
			tap(2),
			tap(3),
		)

		assert.Equal(t, exitcode.Nil, code)
		assert.Equal(t, []int{1, 2, 3, 6, 4, 2}, output)
	})

	t.Run("error throw in capture stage", func(t *testing.T) {
		arguments := []string{"foo", "bar"}
		mDirtyErr := errors.Err(exitcode.RepoDirty, "mock repo is dirty")
		mLockedErr := errors.Err(exitcode.RemoteLocked, "mock remote locked")

		callIns := []runnerCallIn{
			{
				captureReturnErr: nil,
			},
			{
				captureReturnErr: nil,
				bubbleReturnErr: func(e error) error {
					return mLockedErr
				},
			},
			{
				captureReturnErr: mDirtyErr,
			},
			{
				captureReturnErr: nil,
			},
		}

		callOuts, runners := runnerFactory(callIns)

		code := runner.Run(
			arguments,
			runners...,
		)

		assert.Equal(t, len(callIns), len(callOuts))

		t.Run("the downstream runners will not be invoked", func(t *testing.T) {
			assert.Equal(t, 0, callOuts[3].captureCallTimes)
			assert.Equal(t, 0, callOuts[3].bubbleCallTimes)
		})

		t.Run("the bubble stage of failed runner will not be invoked", func(t *testing.T) {
			assert.Equal(t, 0, callOuts[2].bubbleCallTimes)
		})

		t.Run("error was bubbled to upstream in bubble stage", func(t *testing.T) {
			assert.Equal(t, 1, callOuts[0].bubbleCallTimes)
			assert.Equal(t, 1, callOuts[1].bubbleCallTimes)

			t.Run("bubbled error can be changed in runner's bubble stage", func(t *testing.T) {
				assert.Equal(t, mDirtyErr, callOuts[1].bubbleCallErrs[0])
				assert.Equal(t, mLockedErr, callOuts[0].bubbleCallErrs[0])
			})

			t.Run("runner return the code of final bubbled error", func(t *testing.T) {
				assert.Equal(t, errors.GetErrorCode(mLockedErr), code)
			})
		})
	})

	t.Run("recover from the panic of runners", func(t *testing.T) {
		code := runner.Run(
			nil,
			func(args []string) (runner.BubbleTap, error) {
				panic(errors.Err(exitcode.RepoDirty, "mock dirty"))
			},
		)

		assert.Equal(t, exitcode.RepoDirty, code)
	})

	t.Run("skip nil runner", func(t *testing.T) {
		code := runner.Run(
			nil,
			nil,
			func(args []string) (runner.BubbleTap, error) {
				return nil, nil
			},
		)

		assert.Equal(t, exitcode.Nil, code)
	})
}

type runnerCallIn struct {
	captureReturnErr error
	bubbleReturnErr  func(error) error
}

type runnerCallOut struct {
	captureCallTimes int
	captureCallArgs  [][]string

	bubbleCallTimes int
	bubbleCallErrs  []error
}

func runnerFactory(callIns []runnerCallIn) ([]*runnerCallOut, []runner.CaptureTap) {
	callOuts := make([]*runnerCallOut, len(callIns))
	runners := make([]runner.CaptureTap, len(callIns))

	for i, callIn := range callIns {
		callOut := &runnerCallOut{}
		callOuts[i] = callOut

		runners[i] = func(in runnerCallIn, out *runnerCallOut) runner.CaptureTap {
			return func(args []string) (runner.BubbleTap, error) {
				out.captureCallTimes++
				out.captureCallArgs = append(out.captureCallArgs, args)

				return func(e error) error {
					out.bubbleCallTimes++
					out.bubbleCallErrs = append(out.bubbleCallErrs, e)

					if in.bubbleReturnErr != nil {
						return in.bubbleReturnErr(e)
					}
					return e
				}, in.captureReturnErr
			}
		}(callIn, callOut)
	}

	return callOuts, runners
}
