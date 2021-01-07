package repository

type ConfigReader interface {
	GetConfig(keys ...string) (string, error)
}

type ConfigWriter interface {
	SetConfig(key, value string) error
}

type ConfigReadWriter interface {
	ConfigReader
	ConfigWriter
}
