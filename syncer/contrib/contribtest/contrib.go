package contribtest

// MockContrib mock contrib interface
type MockContrib struct {
	StubIn
	StubOut
}

// GetHeadSHA1 mock GetHeadSHA1
func (r *MockContrib) GetHeadSHA1() (string, error) {
	r.GetHeadSHA1CallTimes++

	return r.GetHeadSHA1Return, r.GetHeadSHA1ReturnErr
}

// Sync mock Sync
func (r *MockContrib) Sync(sha1 string, uploads []string, deletes []string) (uploaded []string, deleted []string, err error) {
	r.SyncCallTimes++
	r.SyncCallSHA1 = append(r.SyncCallSHA1, sha1)
	r.SyncCallUploads = append(r.SyncCallUploads, uploads)
	r.SyncCallDeletes = append(r.SyncCallDeletes, deletes)

	err = r.SyncReturnErr

	if r.SyncReturnUploaded != nil {
		uploaded = r.SyncReturnUploaded(uploads)
	}

	if r.SyncReturnDeleted != nil {
		deleted = r.SyncReturnDeleted(deletes)
	}

	return
}
