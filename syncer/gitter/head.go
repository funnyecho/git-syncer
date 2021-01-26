package gitter

// HeadReader repo head reader
type HeadReader interface {
	GetHead() (string, error)
	GetHeadSHA1() (string, error)
}

// HeadWriter repo head writer
type HeadWriter interface {
	PushHead(head string) (string, error)
	PopHead(head string) error
}

// HeadReadWriter repo head reader and writer
type HeadReadWriter = interface {
	HeadReader
	HeadWriter
}
