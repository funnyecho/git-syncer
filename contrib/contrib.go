package contrib

type SyncReq struct {
	SHA1    string
	Uploads []string
	Deletes []string
}

type SyncRes struct {
	SHA1     string
	Uploaded []string
	Deleted  []string
}

type Contrib interface {
	GetHeadSHA1() (string, error)
	Sync(SyncReq) (SyncRes, error)
}
