package log

const (
	VerboseSilent = iota
	VerboseError
	VerboseInfo
	VerboseDebug
)

const LowestVerbose = VerboseDebug

var verbose = VerboseSilent

func UseVerbose() int {
	return min(LowestVerbose, verbose)
}

func WithVerbose(v int) {
	verbose = min(LowestVerbose, v)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}