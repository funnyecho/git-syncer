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

	// MissingArguments missing required arguments
	MissingArguments
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
)

const (
	// ContribForbidden unknown error
	ContribForbidden = 300 + iota

	// ContribHeadNotFound head not found
	ContribHeadNotFound

	// ContribSyncFailed failed to sync
	ContribSyncFailed

	// ContribLocked contrib was locked
	ContribLocked

	// ContribUnlock contrib is unlock
	ContribUnlock

	// ContribInvalidLog invalid log format
	ContribInvalidLog

	// ContribInvalidLock invalid lock format
	ContribInvalidLock
)
