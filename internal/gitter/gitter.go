package gitter

type Gitter interface {
	GetGitVersion() (majorVersion, minorVersion int, err error)
	GetProjectDir() (dir string, err error)
}
