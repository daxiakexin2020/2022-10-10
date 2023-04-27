package bus

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"sync"
)

type store struct {
	client *redis.Client
}

var (
	GStore *store
	sonce  sync.Once
)

func InitStore(c *redis.Client) *store {
	sonce.Do(func() {
		GStore = &store{client: c}
	})
	return GStore
}

func GetStore() (*store, error) {
	if GStore == nil {
		return nil, errors.New("Stroe need init...")
	}
	return GStore, nil
}
