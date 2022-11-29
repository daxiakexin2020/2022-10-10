package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var endpoints = []string{
	"100.69.62.1:3232",
	"100.69.62.2:3232",
	"100.69.62.3:3232",
	"100.69.62.4:3232",
	"100.69.62.5:3232",
	"100.69.62.6:3232",
	"100.69.62.7:3232",
}

func main() {
	testRand()
}

func testRand() {
	indexs := shuffleInner(len(endpoints))
	var maxRetryTimes int = 3
	var checkIndex int = 0
	for i := 0; i < maxRetryTimes; i++ {
		endpoint := endpoints[indexs[checkIndex]]
		//todo 拿到ednpoint，处理业务，比如发送请求，
		err := apiRequest(i)
		if err == nil {
			fmt.Println("endpoint ok", i, endpoint)
			break
		}
		checkIndex++
		fmt.Println("endpoint failed", i, endpoint)
	}
}

// 模拟业务处理
func apiRequest(index int) error {
	if index == 1 {
		return nil
	}
	return errors.New("test error")
}

// shuffle 循环一遍slice，随机生成两个索引，交换这两个索引对应的值，完成洗牌
func shuffle(slice []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice)) //rand.Intn(n)   返回0-n之间的随机数  ，这是伪随机数，都是返回固定的数
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
}

// 改进
func shuffle2(indexes []int) {
	for i := len(indexes); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
	}
}

// Todo go内置的洗牌算法，比上述shuffle方法更科学，更随机，概率更均匀
func shuffleInner(n int) []int {
	rand.Seed(time.Now().UnixNano()) //设置随机种子
	b := rand.Perm(n)
	return b
}
