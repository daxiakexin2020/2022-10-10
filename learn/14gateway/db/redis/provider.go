package redis

import (
	"14gateway/config"
	"fmt"
)

func InitRedis() error {

	rc := config.GetRedisConfig()
	addr := fmt.Sprintf("%s:%d", rc.Addr, rc.Port)
	options := make([]Option, 0)
	if rc.DB != 0 {
		options = append(options, WithDB(rc.DB))
	}
	if rc.Password != "" {
		options = append(options, WithPassword(rc.Password))
	}
	if rc.PoolSize > 0 {
		options = append(options, WithPoolSize(rc.PoolSize))
	}

	_, err := NewRedisClient(addr, options...)
	if err != nil {
		return err
	}
	return nil
}
