package variable

import "github.com/funnyecho/git-syncer/constants"

const (
	// VerboseSilent do not print log
	VerboseSilent = iota
	// VerboseError print error and higher level log
	VerboseError
	// VerboseInfo print info and higher level log
	VerboseInfo
	// VerboseDebug print debug and higher level log
	VerboseDebug
)

// LowestVerbose lowest verbose level
const LowestVerbose = VerboseDebug

// verbose used verbose level
var verbose = VerboseInfo

// WithVerbose update verbose level
func WithVerbose(v constants.Verbose) {
	verbose = min(LowestVerbose, sToI(v))
}

// UseVerbose get used verbose level
func UseVerbose() int {
	return min(LowestVerbose, verbose)
}

// UseVerboseSilent return whether support silent level
func UseVerboseSilent() bool {
	return verbose >= VerboseSilent
}

// UseVerboseError return whether support error level
func UseVerboseError() bool {
	return verbose >= VerboseError
}

// UseVerboseInfo return whether support info level
func UseVerboseInfo() bool {
	return verbose >= VerboseInfo
}

// UseVerboseDebug return whether support debug level
func UseVerboseDebug() bool {
	return verbose >= VerboseDebug
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sToI(v constants.Verbose) int {
	switch v {
	case constants.VerboseSilent:
		return VerboseSilent
	case constants.VerboseError:
		return VerboseError
	case constants.VerboseDebug:
		return VerboseDebug
	case constants.VerboseInfo:
		fallthrough
	default:
		return VerboseInfo
	}
}
