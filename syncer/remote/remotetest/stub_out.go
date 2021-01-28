package remotetest

// StubOut mock calling output
type StubOut struct {
	GetHeadSHA1CallTimes int

	SyncCallTimes   int
	SyncCallSHA1    []string
	SyncCallUploads [][]string
	SyncCallDeletes [][]string
}
