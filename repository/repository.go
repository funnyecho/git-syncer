package repository

type Repository interface {
	ConfigReadWriter
	HeadReadWriter
	Files
}
