package contrib_test

import "github.com/funnyecho/git-syncer/contrib"

type ContribHeadSha1Fetcher func() (string, error)
type ContribSyncHandler func(*contrib.SyncReq) (contrib.SyncRes, error)

type RepoGetHeadSHA1 func() (string, error)
type RepoListAllFiles func() (sha1 string, uploads []string, err error)

type mockContrib struct {
	headSHA1    ContribHeadSha1Fetcher
	syncHandler ContribSyncHandler
}

func (c *mockContrib) GetHeadSHA1() (string, error) {
	if c.headSHA1 == nil {
		return "", nil
	}

	return c.headSHA1()
}

func (c *mockContrib) Sync(req *contrib.SyncReq) (contrib.SyncRes, error) {
	if c.syncHandler == nil {
		return contrib.SyncRes{}, nil
	}

	return c.syncHandler(req)
}
