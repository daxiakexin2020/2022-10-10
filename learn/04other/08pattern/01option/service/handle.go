package service

import "time"

type Redis struct {
	Addr      string
	Port      string
	DB        int
	Expertime time.Time
	Clone     bool
}

type Options func(r *Redis)

func NewClient(addr string, port string, opts ...Options) *Redis {
	r := new(Redis)
	r.Addr = addr
	r.Port = port
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func WithDB(db int) Options {
	return func(r *Redis) {
		r.DB = db
	}
}

func WithExpertime(time time.Time) Options {
	return func(r *Redis) {
		r.Expertime = time
	}
}

func WithClone(clone bool) Options {
	return func(r *Redis) {
		r.Clone = clone
	}
}
