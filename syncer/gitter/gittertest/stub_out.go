package gittertest

// StubOut mock calling output
type StubOut struct {
	GetConfigCallTimes int
	GetConfigCallKeys  []string

	SetConfigCallTimes  int
	SetConfigCallKeys   []string
	SetConfigCallValues []string

	GetHeadCallTimes int

	GetHeadSHA1CallTimes int

	PushHeadCallTimes int
	PushHeadCallHeads []string

	PopHeadCallTimes int
	PopHeadCallHeads []string

	ListTrackedFilesCallTimes int
	ListTrackedFilesCallPaths []string

	ListChangedFilesCallTimes    int
	ListChangedFilesCallPaths    []string
	ListChangedFilesCallBaseSHA1 []string
}
