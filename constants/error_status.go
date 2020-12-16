package constants

const (
	ErrorStatusNoError = iota
	ErrorStatusUnknown
	ErrorStatusUsage
	ErrorStatusMissingArguments
	ErrorStatusUpload
	ErrorStatusRemoteLocked
	ErrorStatusFilesystem
	ErrorStatusGit
)