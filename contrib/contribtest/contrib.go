package contribtest

import "github.com/funnyecho/git-syncer/contrib"

// ContribHeadSha1Fetcher type alias for contrib head fetcher
type ContribHeadSha1Fetcher func() (string, error)

// ContribSyncHandler type alias for contrib sync handler
type ContribSyncHandler func(*contrib.SyncReq) (contrib.SyncRes, error)

// RepoGetHeadSHA1 type alias for get head sha1 from repo
type RepoGetHeadSHA1 func() (string, error)

// RepoListAllFiles type alias for list all files from repo
type RepoListAllFiles func() (sha1 string, uploads []string, err error)

// MockContrib contrib mocking
type MockContrib struct {
	HeadSHA1    ContribHeadSha1Fetcher
	SyncHandler ContribSyncHandler
}

// GetHeadSHA1 implement GetHeadSHA1
func (c *MockContrib) GetHeadSHA1() (string, error) {
	if c.HeadSHA1 == nil {
		return "", nil
	}

	return c.HeadSHA1()
}

// Sync implement Sync
func (c *MockContrib) Sync(req *contrib.SyncReq) (contrib.SyncRes, error) {
	if c.SyncHandler == nil {
		return contrib.SyncRes{}, nil
	}

	return c.SyncHandler(req)
}
