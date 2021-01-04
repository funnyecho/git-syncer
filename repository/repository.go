package repository

type Repository interface {
	GetConfig(keys ...string) (string, error)
	GetHead() (string, error)
	GetHeadSHA1() (string, error)
	GetRepoDir() (string, error)
	GetSyncRoot() (string, error)
	GitVersion() (majorV, minorV int, err error)

	IsDirtyRepository() (bool, error)

	ListAllFiles() ([]string, error)
	ListChangedFiles() (upload []string, delete []string, err error)

	SetHead(head string) error
	SetupTempDir() (string, error)
}

func WithRepository(repo Repository) {
	repository = repo
}

func UseRepository() Repository {
	if repository == nil {
		panic("repository not existed")
	}

	return repository
}

var repository Repository
