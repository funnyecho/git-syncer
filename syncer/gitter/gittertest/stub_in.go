package gittertest

// StubIn mock calling source
type StubIn struct {
	GetConfigReturn    string
	GetConfigReturnErr error
	GetConfigReturnFn  func(string) (string, error)

	SetConfigReturnErr error

	GetHeadReturn    string
	GetHeadReturnErr error

	GetHeadSHA1Return    string
	GetHeadSHA1ReturnErr error

	PushHeadReturn    string
	PushHeadReturnErr error

	PopHeadReturnErr error

	ListTrackedFilesReturn    []string
	ListTrackedFilesReturnErr error

	ListChangedFilesReturnAM  []string
	ListChangedFilesReturnD   []string
	ListChangedFilesReturnErr error
}
