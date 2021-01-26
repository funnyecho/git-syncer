package exitcode

const (
	// Nil mean successfully
	Nil = iota
)

const (
	// Unknown unknown error
	Unknown = 100 + iota

	// Usage invalid usage
	Usage

	// MissingArguments missing required arguments for command
	MissingArguments

	// InvalidParams invalid func params
	InvalidParams

	// InvalidRunnerDependency means invalid runner dependency
	InvalidRunnerDependency
)

const (
	// RepoUnknown unknown error
	RepoUnknown = 200 + iota

	// RepoCheckoutFailed failed to checkout to head
	RepoCheckoutFailed

	// RepoDirty repository is dirty
	RepoDirty

	// RepoHeadNotFound repository head not found
	RepoHeadNotFound

	// RepoListFilesFailed list repo files failed
	RepoListFilesFailed

	// RepoDiffBaseNotFound basic commit to diff is empty
	RepoDiffBaseNotFound

	// RepoConfigNotFound config not found
	RepoConfigNotFound

	// RepoInvalidGitVersion git version not supported
	RepoInvalidGitVersion
)

const (
	// RemoteForbidden unknown error
	RemoteForbidden = 300 + iota

	// RemoteHeadNotFound head not found
	RemoteHeadNotFound

	// RemoteSyncFailed failed to sync
	RemoteSyncFailed

	// RemoteLocked Remote was locked
	RemoteLocked

	// RemoteUnlock Remote is unlock
	RemoteUnlock

	// RemoteInvalidLog invalid log format
	RemoteInvalidLog

	// RemoteInvalidLock invalid lock format
	RemoteInvalidLock
)
