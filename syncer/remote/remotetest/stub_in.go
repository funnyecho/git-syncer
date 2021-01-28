package remotetest

// StubIn mock calling source
type StubIn struct {
	GetHeadSHA1Return    string
	GetHeadSHA1ReturnErr error

	SyncReturnUploaded func([]string) []string
	SyncReturnDeleted  func([]string) []string
	SyncReturnErr      error
}
