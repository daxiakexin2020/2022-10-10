package service

import (
	"context"
	"fmt"
	"github.com/agext/uuid"
	"time"
)

var (
	Total    int
	GetCount int
)

func (c *Rclient) Get(key string) (string, error) {
	return c.client.Get(context.Background(), key).Result()
}

func (c *Rclient) Set(key string, val interface{}, expiration time.Duration) error {
	return c.client.Set(context.Background(), key, val, expiration).Err()
}

func (r *Rclient) HSet(key string, field string, val interface{}) (int64, error) {
	return r.client.HSet(context.Background(), key, field, val).Result()
}

func (r *Rclient) HDel(key string, field string) (int64, error) {
	return r.client.HDel(context.Background(), key, field).Result()
}

func (r *Rclient) HGet(key string, field string) (string, error) {
	return r.client.HGet(context.Background(), key, field).Result()
}

func (r *Rclient) Del(key string) (int64, error) {
	return r.client.Del(context.Background(), key).Result()
}

func (rc *Rclient) DispersedLock(key string) (bool, error) {

	var ctx, cancel = context.WithCancel(context.Background())

	defer func() {
		cancel()
	}()

	//1  不存在此key，设置成功，同时设置过期时间
	value := string(uuid.NodeId())
	ifSetOk, err := rc.client.SetNX(ctx, key, value, 1*time.Millisecond).Result()
	if !ifSetOk {
		return ifSetOk, err
	}

	//2  自动续期
	go rc.WatchDog(ctx, key, 2*time.Second, "test")

	//todo 模拟处理业务
	GetCount++
	Total++

	//3 业务完成，原子性操作，使用lua脚本 删除锁
	lua := `
	-- 如果当前值与锁值一致,删除key
	if redis.call('GET', KEYS[1]) == ARGV[1] then
		return redis.call('DEL', KEYS[1])
	else
		return 0
	end
`
	val, err := rc.client.Eval(context.Background(), lua, []string{key}, value).Result()
	if err != nil {
		panic(err.Error())
	}

	return val == int64(1), nil
}

// 自动续期
func (rc *Rclient) WatchDog(ctx context.Context, key string, expiration time.Duration, tag string) {
	for {
		select {
		// 业务完成
		case <-ctx.Done():
			fmt.Printf("%s任务完成,关闭%s的自动续期\n", tag, key)
			return
			// 业务未完成
		default:
			// 自动续期
			rc.client.PExpire(ctx, key, expiration)
			// 继续等待
			time.Sleep(expiration / 2)
		}
	}
}
