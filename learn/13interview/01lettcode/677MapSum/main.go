package main

import (
	"fmt"
	"strings"
)

func main() {
	ms := Constructor()
	ms.Insert("key1", 1)
	ms.Insert("key2", 2)
	ms.Insert("key3", 3)
	ms.Insert("nkey1", 1)

	fmt.Printf("count=%d\n", ms.Sum("key"))
}

type MapSum struct {
	data map[string]int
}

func Constructor() MapSum {
	return MapSum{
		data: map[string]int{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.data[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	var count int
	for k, v := range this.data {
		if strings.HasPrefix(k, prefix) {
			count += v
		}
	}
	return count
}
