package contrib

type Contrib interface {
	CheckAccessible() error
	GetHeadSHA1(remote string) (string, error)
	Lock() error
	Unlock() error
	RLock() error
	RUnlock() error
}

func WithContrib(c Contrib) {
	contrib = c
}

func UseContrib() Contrib {
	if contrib == nil {
		panic("contrib not existed")
	}
	return contrib
}

var contrib Contrib
