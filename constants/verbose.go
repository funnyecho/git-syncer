package constants

// Verbose const type
type Verbose = string

const (
	// VerboseSilent do not print log
	VerboseSilent = Verbose("silent")
	// VerboseError print error and higher level log
	VerboseError = Verbose("error")
	// VerboseInfo print info and higher level log
	VerboseInfo = Verbose("info")
	// VerboseDebug print debug and higher level log
	VerboseDebug = Verbose("debug")
)
