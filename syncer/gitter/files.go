package gitter

// Files interface to list repo files
type Files interface {
	ListTrackedFiles(path string) (uploads []string, err error)
	ListChangedFiles(path string, baseSha1 string) (amFiles []string, dFiles []string, err error)
}
