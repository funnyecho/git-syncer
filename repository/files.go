package repository

type Files interface {
	ListAllFiles() (sha1 string, uploads []string, err error)
	ListChangedFiles(baseSha1 string) (sha1 string, uploads []string, deletes []string, err error)
}
