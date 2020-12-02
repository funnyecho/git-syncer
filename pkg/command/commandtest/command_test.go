package commandtest_test

import (
	"fmt"
	"github.com/funnyecho/git-syncer/pkg/command/commandtest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTestingCommander(t *testing.T) {
	t.Run("commander options", func(t *testing.T) {
		t.Run("shall panic with invalid helper name", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					assert.Fail(t, "shall panic when helper name not start with `Test`")
				}
			}()

			commandtest.NewTestingCommander(
				commandtest.WithCommanderHelperName("NotStartWithTest"),
			)
		})

		helperName := "TestCommanderOptions"

		envFactoryCallTimes := 0
		envFactoryCallCmdName := ""

		commander := commandtest.NewTestingCommander(
			commandtest.WithCommanderHelperName(helperName),
			commandtest.WithCommanderEnvs(func(prevEnvs []string, name string, args ...string) []string {
				envFactoryCallCmdName = name
				envFactoryCallTimes += 1
				return prevEnvs
			}),
		)

		cmd := commander("foo")
		assert.Equal(t, 1, envFactoryCallTimes)
		assert.Equal(t, "foo", envFactoryCallCmdName)

		if len(cmd.Args) >= 2 {
			assert.Equal(t, fmt.Sprintf("-test.run=%s", helperName), cmd.Args[1])
		} else {
			assert.Fail(t, "testing command shall has at least two arguments")
		}

		cmd = commander("bar")
		assert.Equal(t, 2, envFactoryCallTimes)
		assert.Equal(t, "bar", envFactoryCallCmdName)
	})
}
