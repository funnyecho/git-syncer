package contrib

type Contrib interface {

}

var NoopContrib = &noopContrib{}

type noopContrib struct {

}
