package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	RClient *rclient
	once    sync.Once
)

type rclient struct {
	config *redis.Options
	Client *redis.Client
}

type Option func(c *rclient)

func WithDB(db int) Option {
	return func(r *rclient) {
		r.config.DB = db
	}
}

func WithPassword(password string) Option {
	return func(r *rclient) {
		r.config.Password = password
	}
}

func (r *rclient) apply(opts []Option) {
	for _, opt := range opts {
		opt(r)
	}
}

func NewClient(addr string, opts ...Option) (*rclient, error) {

	var err error
	once.Do(func() {
		rc := &rclient{config: &redis.Options{
			Addr: addr,
		}}
		rc.apply(opts)
		cli := redis.NewClient(rc.config)
		if err = cli.Ping(context.Background()).Err(); err == nil {
			rc.Client = cli
			RClient = rc
		}
	})
	return RClient, err
}
