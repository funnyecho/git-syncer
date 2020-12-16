package contrib

type Contrib interface {

}

func WithContrib(c Contrib) {
	contrib = c
}

func UseContrib() Contrib {
	return contrib
}

var contrib Contrib = &noopContrib{}

type noopContrib struct {}
