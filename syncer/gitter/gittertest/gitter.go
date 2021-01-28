package gittertest

// MockGitter mock gitter interface
type MockGitter struct {
	StubIn
	StubOut
}

// GetConfig mock GetConfig
func (g *MockGitter) GetConfig(key string) (string, error) {
	g.GetConfigCallTimes++
	g.GetConfigCallKeys = append(g.GetConfigCallKeys, key)

	if g.GetConfigReturnFn != nil {
		return g.GetConfigReturnFn(key)
	}

	return g.GetConfigReturn, g.GetConfigReturnErr
}

// SetConfig mock SetConfig
func (g *MockGitter) SetConfig(key string, value string) error {
	g.SetConfigCallTimes++
	g.SetConfigCallKeys = append(g.SetConfigCallKeys, key)
	g.SetConfigCallValues = append(g.SetConfigCallValues, value)

	return g.SetConfigReturnErr
}

// GetHead mock GetHead
func (g *MockGitter) GetHead() (string, error) {
	g.GetHeadCallTimes++

	return g.GetHeadReturn, g.GetHeadReturnErr
}

// GetHeadSHA1 mock GetHeadSHA1
func (g *MockGitter) GetHeadSHA1() (string, error) {
	g.GetHeadSHA1CallTimes++

	return g.GetHeadSHA1Return, g.GetHeadSHA1ReturnErr
}

// PushHead mock PushHead
func (g *MockGitter) PushHead(head string) (string, error) {
	g.PushHeadCallTimes++
	g.PushHeadCallHeads = append(g.PushHeadCallHeads, head)

	return g.PushHeadReturn, g.PushHeadReturnErr
}

// PopHead mock PopHead
func (g *MockGitter) PopHead(head string) error {
	g.PopHeadCallTimes++
	g.PopHeadCallHeads = append(g.PopHeadCallHeads, head)

	return g.PopHeadReturnErr
}

// ListTrackedFiles mock ListTrackedFiles
func (g *MockGitter) ListTrackedFiles(path string) (uploads []string, err error) {
	g.ListTrackedFilesCallTimes++
	g.ListTrackedFilesCallPaths = append(g.ListTrackedFilesCallPaths, path)

	return g.ListTrackedFilesReturn, g.ListTrackedFilesReturnErr
}

// ListChangedFiles mock ListChangedFiles
func (g *MockGitter) ListChangedFiles(path string, baseSHA1 string) (amFiles []string, dFiles []string, err error) {
	g.ListChangedFilesCallTimes++
	g.ListChangedFilesCallPaths = append(g.ListChangedFilesCallPaths, path)
	g.ListChangedFilesCallBaseSHA1 = append(g.ListChangedFilesCallBaseSHA1, baseSHA1)

	return g.ListChangedFilesReturnAM, g.ListChangedFilesReturnD, g.ListChangedFilesReturnErr
}
