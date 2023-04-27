package service

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Rclient struct {
	config *redis.Options
	client *redis.Client
}

type Option func(c *Rclient)

func WithDB(db int) Option {
	return func(r *Rclient) {
		r.config.DB = db
	}
}

func WithPassword(password string) Option {
	return func(r *Rclient) {
		r.config.Password = password
	}
}

func (r *Rclient) apply(opts []Option) {
	for _, opt := range opts {
		opt(r)
	}
}

func NewClient(addr string, opts ...Option) (*Rclient, error) {

	rc := &Rclient{config: &redis.Options{
		Addr: addr,
	}}
	rc.apply(opts)
	cli := redis.NewClient(rc.config)
	if err := cli.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	rc.client = cli
	return rc, nil
}
