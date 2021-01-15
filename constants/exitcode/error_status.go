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
)

const (
	// ContribUnknown unknown error
	ContribUnknown = 300 + iota

	// ContribForbidden forbidden
	ContribForbidden

	// ContribHeadNotFound head not found
	ContribHeadNotFound

	// ContribSyncFailed failed to sync
	ContribSyncFailed
)
