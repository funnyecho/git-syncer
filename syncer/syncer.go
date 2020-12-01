package syncer

type Syncer interface {

}

func New() Syncer {
	return &syncer{}
}

type syncer struct {

}
