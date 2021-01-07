package repository

type HeadReader interface {
	GetHead() (string, error)
	GetHeadSHA1() (string, error)
}

type HeadWriter interface {
	PushHead(head string) (string, error)
	PopHead(head string) error
}

type HeadReadWriter interface {
	HeadReader
	HeadWriter
}
