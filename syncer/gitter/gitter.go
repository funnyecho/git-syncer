package gitter

// Gitter interface for gitter
type Gitter interface {
	ConfigReadWriter
	HeadReadWriter
	Files
}
