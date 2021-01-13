package repository

// ConfigReader config reader
type ConfigReader interface {
	GetConfig(key string) (string, error)
}

// ConfigWriter config writer
type ConfigWriter interface {
	SetConfig(key, value string) error
}

// ConfigReadWriter config reader & writer
type ConfigReadWriter interface {
	ConfigReader
	ConfigWriter
}
