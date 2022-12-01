package service

import (
	"context"
	"fmt"
	"github.com/agext/uuid"
	"time"
)

func (c *rclient) Get(key string) (string, error) {
	return c.Client.Get(context.Background(), key).Result()
}

func (c *rclient) Set(key string, val interface{}, expiration time.Duration) error {
	return c.Client.Set(context.Background(), key, val, expiration).Err()
}

var Total int
var GetCount int

func (rc *rclient) DispersedLock(key string) (bool, error) {

	var ctx, cancel = context.WithCancel(context.Background())

	defer func() {
		cancel()
	}()

	//1  不存在此key，设置成功，同时设置过期时间
	value := string(uuid.NodeId())
	ifSetOk, err := rc.Client.SetNX(ctx, key, value, 1*time.Millisecond).Result()
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
	val, err := rc.Client.Eval(context.Background(), lua, []string{key}, value).Result()
	if err != nil {
		panic(err.Error())
	}

	return val == int64(1), nil
}

// 自动续期
func (rc *rclient) WatchDog(ctx context.Context, key string, expiration time.Duration, tag string) {
	for {
		select {
		// 业务完成
		case <-ctx.Done():
			fmt.Printf("%s任务完成,关闭%s的自动续期\n", tag, key)
			return
			// 业务未完成
		default:
			// 自动续期
			rc.Client.PExpire(ctx, key, expiration)
			// 继续等待
			time.Sleep(expiration / 2)
		}
	}
}
