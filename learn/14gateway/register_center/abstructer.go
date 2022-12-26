package register_center

import "context"

type Register interface {
	Name() string
	Get(ctx context.Context, key string) ([]string, error)
	Put(ctx context.Context, key string, val string, options ...interface{}) error
	Delete(ctx context.Context, key string) error
}
