package model

type ServerRepo interface {
	DefaultServerRun() error
	Run(addr string) error
	Stop() error
}
