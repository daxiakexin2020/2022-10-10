package items

import (
	"22go_redis/server/construct"
	"22go_redis/utils"
	"fmt"
	"strconv"
	"sync"
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

	m := make(map[interface{}]int64)
	m[10] = 1
	m["20"] = 2
	m[70] = 7
	m[50] = 5
	m[60] = 6
	GORedis.Zadd("test02", m)
	zrevrange, i := GORedis.Zrevrange("test02", -100000, 100000)
	fmt.Println("Zrevrange:", zrevrange, i)

	GORedis.Zincrby("test02", 10, -100)
	zrevrange2, i := GORedis.Zrevrange("test02", -100000, 10000)
	fmt.Println("Zrevrange2:", zrevrange2, i)
}

func TestB12(t *testing.T) {
	m := make(map[interface{}]int64)
	m[1] = 1
	m["20"] = 2
	m[70] = 7
	m[50] = 5
	m[60] = 6
	GORedis.Zadd("test02", m)
	zrevrange, i := GORedis.Zrevrange("test02", -100000, 100000)
	fmt.Println("Zrevrange:", zrevrange, i)

	m2 := make(map[interface{}]int64)
	m2[1] = 100
	m2[40] = 400
	m2["20"] = 2
	m2[70] = 7
	m2[50] = 5
	m2[60] = 6
	GORedis.Zadd("test02", m2)
	zrevrange2, i2 := GORedis.Zrevrange("test02", -100000, 100000)
	fmt.Println("Zrevrange2:", zrevrange2, i2)

	zrank, err := GORedis.Zrank("test02", "20")
	fmt.Println("zrank:", zrank, err)

	zrem := GORedis.Zrem("test02", 1, "20", 40, 70, 50, 60, 4444)
	fmt.Println("Zrem:", zrem)

	zrevrange3, i3 := GORedis.Zrevrange("test02", -100000, 100000)
	fmt.Println("Zrevrange3:", zrevrange3, i3)

	m3 := make(map[interface{}]int64)
	m3[1] = 100
	m3[40] = 400
	m3["20"] = 2
	m3[70] = 7
	m3[50] = 5
	m3[60] = 6
	GORedis.Zadd("test02", m3)
	zrevrange4, i4 := GORedis.Zrevrange("test02", -100000, 100000)
	fmt.Println("Zrevrange4:", zrevrange4, i4)

}

func TestB13(t *testing.T) {
	clist := construct.NewClist("1")
	clist.ShowFromHead()
	clist.ShowFromTail()
}

func TestB14(t *testing.T) {
	clist := construct.NewClist("10")
	clist.AddToHead(9)
	clist.AddToHead(8)
	clist.ShowFromHead()
	clist.ShowFromTail()

	size := clist.Size()
	fmt.Println("size:", size)
}

func TestB15(t *testing.T) {
	clist := construct.NewClist("10")
	clist.AddToHead(9)
	clist.AddToHead(8)
	clist.ShowFromHead()

	rn1 := clist.RemoveHead()
	rn2 := clist.RemoveHead()

	fmt.Println("rn1:", rn1.GetVal())
	fmt.Println("rn2:", rn2.GetVal())

	clist.ShowFromHead()
	clist.ShowFromTail()

	tail := clist.RemoveTail()
	fmt.Println("tail1:", tail)

	tail2 := clist.RemoveTail()
	fmt.Println("tail2:", tail2)

	size := clist.Size()
	fmt.Println("size:", size)
}

func TestB16(t *testing.T) {

	clist := construct.NewClist("10")
	clist.AddToTail(9)
	clist.AddToTail(8)
	clist.RemoveTail()
	clist.ShowFromTail()

	size := clist.Size()
	fmt.Println("size:", size)
}

func TestB17(t *testing.T) {

	clist := construct.NewClist("10")
	clist.AddToHead(9)
	clist.AddToHead(8)
	clist.ShowFromHead()
	clist.ShowFromHead()

	get := clist.Get(1)
	fmt.Println("Get:", get)
}

func TestB18(t *testing.T) {

	clist := construct.NewClist("10")
	clist.AddToHead(9)
	clist.AddToHead(8)
	clist.ShowFromHead()

	remove := clist.Remove(1)
	fmt.Println("Remove:", remove)

	clist.ShowFromTail()
	clist.ShowFromHead()
}

func TestB19(t *testing.T) {
	GORedis.Lpush("test02", 1)
	GORedis.Lpush("test02", 2, 3)

	lpop := GORedis.Lpop("test02")
	fmt.Println("val:", lpop)

	l := GORedis.Llen("test02")
	fmt.Println("llen:", l)
}

func TestB20(t *testing.T) {
	GORedis.Lpush("test02", 1, 2, 3, 4, 5, 1)
	//lrange := GORedis.Lrange("test02", 11, 5)
	//fmt.Println("lrange:", lrange)

	lrem := GORedis.Lrem("test02", 1, 1)
	fmt.Println("Lrem:", lrem)
	lrange := GORedis.Lrange("test02", 0, 10)
	fmt.Println("Lrange:", lrange)

	vInterface := GORedis.CGOGet("test02")
	clist := vInterface.(*construct.Clist)
	clist.ShowFromHead()
	clist.ShowFromTail()
}

func TestB21(t *testing.T) {
	GORedis.Lpush("test02", 1, 2, 3, 4, 5)
}

func TestB22(t *testing.T) {
	GORedis.Lpush("test02", 1, 2, 3, 4, 5)
	lset := GORedis.Lset("test02", 0, 40)
	fmt.Println("Lset:", lset)
	get := GORedis.CGOGet("dddd")
	fmt.Println("CGOGet:", get)
}

func TestB23(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			field := strconv.Itoa(i)
			GORedis.Hset("test02", field, i)
		}(i)
	}
	wg.Wait()
	fmt.Println("hset over", GORedis.Hlen("test02"))
	hget := GORedis.Hget("test02", "2")
	fmt.Println("hget:", hget)
}
