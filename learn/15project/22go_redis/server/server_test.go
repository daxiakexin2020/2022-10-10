package server

import (
	"22go_redis/utils"
	"fmt"
	"testing"
	"time"
)

func TestB(t *testing.T) {
	str := "1010mb"
	parse, s, err := Parse(str)
	fmt.Println("ss:", parse, s, err)
}

func TestB02(t *testing.T) {
	gredis, err := NewGredis("10MB")
	fmt.Println("f", gredis, err)
	fmt.Println("go:", GORedis)

	gredis.Set("test", 1, 1000)
	get, b := gredis.Get("test")
	fmt.Println("get:", get, b)

	ttl := gredis.Ttl("test")
	fmt.Println("ttl:", ttl)
}

func TestB03(t *testing.T) {
	GORedis.Set("test", "testval", 10)
	//get, _ := GORedis.Get("test")
	//fmt.Println("get:", get)
	//keys := GORedis.Keys()
	//fmt.Println("keys:", keys)

	//strlen := GORedis.STRLEN("test")
	//fmt.Println("strlen:", strlen, len("testval"))

	for i := 0; i < 20; i++ {
		ttl := GORedis.Ttl("test")
		fmt.Println("ttl:", ttl)
		time.Sleep(time.Second * 1)
		if i == 5 {
			GORedis.EXPIRE("test", 10)
		}
	}
}

func TestB04(t *testing.T) {
	GORedis.Set("test01", []int{1, 2}, 10)
	get, _ := GORedis.Get("test01")
	fmt.Println("get:", get.([]int))
}

func TestB05(t *testing.T) {
	GORedis.Set("test01", 1, 100)
	get, _ := GORedis.Get("test01")
	fmt.Println("get:", get)

	GORedis.EXPIRE("test01", 20)
	ttl := GORedis.Ttl("test01")
	fmt.Println("ttl:", ttl)

	s := GORedis.Type("test01")
	fmt.Println("type:", s)
}

func TestB06(t *testing.T) {
	add1 := GORedis.Sadd("test01", 1, 100)
	add2 := GORedis.Sadd("test01", 1, 100)
	fmt.Println("add:", add1, add2)
	i := GORedis.Smembers("test01")
	fmt.Println("Smembers:", i)
	s := GORedis.Type("test01")
	fmt.Println("type:", s)
}

func TestB07(t *testing.T) {
	GORedis.Set("test02", int64(1), 100)
	i, err := GORedis.STRLEN("test02")
	fmt.Println("STRLEN:", i, err)
	for i := 0; i < 2000; i++ {
		go func() {
			err = GORedis.Decr("test02")
			fmt.Println("Decr:", err)
			get, b := GORedis.Get("test02")
			fmt.Println("get:", get, b)
		}()
	}
	time.Sleep(time.Second)
	get, b := GORedis.Get("test02")
	fmt.Println("get111111:", get, b)
}

func TestB08(t *testing.T) {
	GORedis.Sadd("test02", 1, 2, 3, 1)
	fmt.Println("GORedis.Smembers(\"test02\"):", GORedis.Smembers("test02"))
	smove := GORedis.Smove("test02", "test01", 3)
	fmt.Println("smove:", smove)
	fmt.Println("GORedis.Smembers(\"test02\"):", GORedis.Smembers("test02"))
	fmt.Println("GORedis.Smembers(\"test01\"):", GORedis.Smembers("test01"))
	fmt.Println("GORedis.Scard(\"test01\"):", GORedis.Scard("test02"))

	ints := utils.RandInts(10, 10)
	fmt.Println("index:", ints)
}

func TestB09(t *testing.T) {
	GORedis.Sadd("test02", 1, 2, 3, 4, 5, 6, 7)
	spop := GORedis.Srandmember("test02", 2)
	fmt.Println("Spop:", spop)

	smembers := GORedis.Smembers("test02")
	fmt.Println("Smembers:", smembers)
}

func TestB10(t *testing.T) {
	GORedis.Sadd("test02", 1, 2, 3, 4, 5, 6, 7)
	spop := GORedis.Srem("test02", 7)
	fmt.Println("Srem:", spop)

	smembers := GORedis.Smembers("test02")
	fmt.Println("Smembers:", smembers)
}

func TestB11(t *testing.T) {
	m := make(map[int64]interface{})
	m[10] = "1"
	m[20] = "2"
	m[30] = 3
	m[60] = 6
	m[50] = 5
	GORedis.Zadd("test02", m)
	zrevrange, i := GORedis.Zrevrange("test02", 10, 100)
	fmt.Println("Zrevrange:", zrevrange, i)

	m2 := make(map[int64]interface{})
	m2[100] = "1"
	m2[20] = "2"
	m2[30] = 3
	m2[70] = 6
	m2[50] = 5
	GORedis.Zadd("test02", m2)
	zrevrange2, i2 := GORedis.Zrevrange("test02", 10, 100)
	fmt.Println("Zrevrange2:", zrevrange2, i2)
}
