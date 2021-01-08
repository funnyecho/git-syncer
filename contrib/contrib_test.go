package contrib_test

import "github.com/funnyecho/git-syncer/contrib"

type ContribHeadSha1Fetcher func() (string, error)
type ContribSyncHandler func(*contrib.SyncReq) (contrib.SyncRes, error)

type RepoListAllFiles func() (sha1 string, uploads []string, err error)
