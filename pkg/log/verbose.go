package log

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

// UseVerbose get used verbose level
func UseVerbose() int {
	return min(LowestVerbose, verbose)
}

// WithVerbose update verbose level
func WithVerbose(v int) {
	verbose = min(LowestVerbose, v)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
