package commandtest

import (
	"fmt"
	"os"
	"strconv"
)

const (
	GO_WANT_HELPER_PROCESS = "__GO_WANT_HELPER_PROCESS__"
	COMMAND_TIMES          = "__COMMAND_TIMES__"
)

func EnvWithGoWantHelperProcess() string {
	return fmt.Sprintf("%s=1", GO_WANT_HELPER_PROCESS)
}

func EnvUseGoWantHelperProcess() string {
	return os.Getenv(GO_WANT_HELPER_PROCESS)
}

func EnvWithCommandTimes(times int) string {
	return fmt.Sprintf("%s=%d", COMMAND_TIMES, times)
}

func EnvUseCommandTimes() int {
	raw := os.Getenv(COMMAND_TIMES)
	if raw == "" {
		return 0
	}

	times, timesErr := strconv.Atoi(raw)
	if timesErr != nil {
		return 0
	}

	return times
}
