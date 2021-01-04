package exitcode

const (
	Nil = iota
	Unknown
	Usage
	MissingArguments
	Upload
	RemoteForbidden
	Filesystem
	Git
)
