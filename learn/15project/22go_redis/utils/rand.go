package utils

import (
	"math/rand"
	"time"
)

func RandInts(l int, n int) []int {
	if n > l {
		return nil
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var ret []int
	flag := make(map[int]struct{})
	for len(ret) < n {
		index := r.Intn(l)
		if _, ok := flag[index]; !ok {
			ret = append(ret, index)
			flag[index] = struct{}{}
		}
	}
	return ret
}
