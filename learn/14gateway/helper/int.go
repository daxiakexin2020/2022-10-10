package helper

import (
	"errors"
	"math/rand"
	"time"
)

// return 【0,len)
func RandInt(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("随机数长度必须大于0")
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n), nil
}
