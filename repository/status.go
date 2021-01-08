package repository

type Status interface {
	IsDirtyRepository() (bool, error)
}
