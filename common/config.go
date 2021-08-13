package common

type Config interface {
	GetAddress() string
	GetPort() int
	GetDBFileName() string
}
