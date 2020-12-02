package commandtest

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EnvGoWantHelperProcess = "__GO_WANT_HELPER_PROCESS__"
	EnvCommandTimes        = "__GO_COMMAND_TIMES__"
)

func EnvWithGoWantHelperProcess() string {
	return fmt.Sprintf("%s=1", EnvGoWantHelperProcess)
}

func EnvUseGoWantHelperProcess() bool {
	return os.Getenv(EnvGoWantHelperProcess) == "1"
}

func EnvWithCommandTimes(times int) string {
	return fmt.Sprintf("%s=%d", EnvCommandTimes, times)
}

func EnvUseCommandTimes() int {
	raw := os.Getenv(EnvCommandTimes)
	if raw == "" {
		return 0
	}

	times, timesErr := strconv.Atoi(raw)
	if timesErr != nil {
		return 0
	}

	return times
}
