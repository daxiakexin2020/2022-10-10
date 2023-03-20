package file

import (
	"20red_police/internal/data"
)

type WR interface {
	Read() error
	Write(data interface{}) error
	Close() error
	data.Class
}

func Register(wr ...data.Class) error {
	return data.GclassTree().Register(wr...)
}
