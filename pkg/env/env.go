package env

import (
	"os"
)

const RunMode = "RUN_MODE"

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

var modeName = "debug"

func init() {
	mode := os.Getenv(RunMode)
	if mode != "" {
		SetMode(mode)
	}
}

func SetMode(mode string) {
	modeName = mode
}

func Mode() string {
	return modeName
}

func IsDebugging() bool {
	return modeName == DebugMode
}

func IsTesting() bool {
	return modeName == TestMode
}

func IsReleasing() bool {
	return modeName == ReleaseMode
}
