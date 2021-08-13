package common

type Authority interface {
	GetConfig() Config
	GetDB() Database
	Shutdown()
}
