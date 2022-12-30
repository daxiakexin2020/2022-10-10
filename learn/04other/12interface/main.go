package main

import "fmt"

func main() {
	test()
}

type Getter interface {
	Get(key string) string
}

type GetterFunc func(key string) string

func (f GetterFunc) Get(key string) string {
	return f(key)
}

func GetFromSource(getter Getter, key string) string {
	return getter.Get(key)
}

type DB struct {
	name string
	typ  string
}

func (db *DB) Get(key string) string {
	return "方式三" + key
}

// 使用
func test() {

	//使用1，直接传函数，GetterFunc类型的函数作为参数，传入
	key1 := GetFromSource(GetterFunc(func(key string) string {
		return "方式一" + key
	}), "test")

	//使用2，普通函数传入 将 test2 强制类型转换为 GetterFunc，GetterFunc 实现了接口 Getter，是一个合法参数。这种方式适用于逻辑较为简单的场景。
	key2 := GetFromSource(GetterFunc(test2), "test")

	//使用3，传入实现了接口的结构体
	key3 := GetFromSource(new(DB), "test")

	fmt.Printf("方式1：=%s\n方式2：=%s\n方式3：=%s\n", key1, key2, key3)

}

func test2(key string) string {
	return "方式二" + key
}
