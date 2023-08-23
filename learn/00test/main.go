package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/oliveagle/jsonpath"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"text/template"
	"time"
)

type Code uint32

type Msg string

type Info struct {
	Name string `json:"name"`
	sync.Mutex
}

type List struct {
	Id   string `json:"id"`
	Data []Item `json:"data"`
}

type Item struct {
	Age    int `json:"age"`
	Salary int `json:"salary"`
}

var m = map[string]*Info{}

func test20() *Info {
	m["zz"] = &Info{Name: "kx"}
	return m["zz"]
}

var (
	AppKey    = "344055c59b8a464c8123bd3d40d68474"
	AppSecret = "7edcc56799c34447b4ed38e08dc29e40"
	RequestId = "0f8424e6-f3cc-4a64-a747-d835ae824e1c"
	Created   = "2023-07-05T18:26:42.759Z"
)

func main() {
	GenSignature()
}
func GenSignature() {
	str := formatString(RequestId, Created, AppSecret)
	fmt.Println(genSignature(str))
}

func genSignature(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hs := hash.Sum(nil)
	base64String := base64.StdEncoding.EncodeToString(hs)
	return url.QueryEscape(base64String)
}

func formatString(requestId, created, appSecret string) string {
	return requestId + created + appSecret
}

func main1() {

	jsonStr := "{\n  \"FO_EXPENSEBILL\": {\n    \"ADVANCEMONEY\": 0,\n    \"BILLDATE\": 0,\n    \"CONTRACTNUM\": \"\",\n    \"DEPARTMENTID\": \"\",\n    \"INVOICEMONEY\": 0,\n    \"ISQUOTECONTRACT\": \"\",\n    \"LOANMONEY\": 0,\n    \"MAINBODY\": \"中国铁建国际集团有限公司本级（生产验证）\",\n    \"MEMO\": \"许江报销\",\n    \"OPERATORDEPARTMENT\": \"财务部\",\n    \"PAYMONEY\": 0,\n    \"PAYTYPECODE\": \"010银行支付\",\n    \"PERSONTAX\": \"\",\n    \"RUNNINGMONEY\": 0,\n    \"STAFFCODE\": \"许江\",\n    \"UNITCODE\": \"财务部\"\n  },\n  \"FO_EXPENSEBILL_INVVATITEM\": [\n    {\n      \"BILLDETAITYPE\": \"\",\n      \"BILLTYPE\": \"增值税普通发票\",\n      \"CONSUMPTIONTYPE\": \"差旅费\",\n      \"CURRMONEY\": 0,\n      \"INVOICECHECKCODE\": \"82322969860869404193\",\n      \"INVOICECODE\": \"034022000104\",\n      \"INVOICEDATE\": 1623945600,\n      \"INVOICENUM\": \"02802144\",\n      \"MONEYTOTAL\": 904,\n      \"NOTETHEINVOICE\": \"桥喜家果专用章91340200MA2NTBJD7W\",\n      \"PURNAME\": \"芜湖市人民防空办公室\",\n      \"PURTAXNO\": \"113402000030106774\",\n      \"SUPPLIERNAME\": \"芜湖昶隆办公设备销售有限责任公司\",\n      \"SUPPLIERTAXNO\": \"91340200MA2NTBJD7W\",\n      \"TAXMONEY\": 0,\n      \"TAXRATE\": \"0.13\",\n      \"TICKETNUMBER\": 0\n    }\n  ],\n  \"FO_EXPENSEBILL_ITEM\": [\n    {\n      \"BILLDETAITYPE\": \"增值税普通发票\",\n      \"BUSINESSSUBCLASS\": \"\",\n      \"CONSUMPTIONTYPE\": \"差旅费\",\n      \"DEPARTMENTCODE\": \"\",\n      \"EXPENSEATTRIBUTE\": \"\",\n      \"TAXMONEY_N\": 104.61,\n      \"TAXRATE\": 0.13,\n      \"TICKETMONEY\": 799.99\n    }\n  ]\n}"
	jsonByte := []byte(jsonStr)

	var obj interface{}
	path := "$.FO_EXPENSEBILL.MAINBODY"
	err := json.Unmarshal(jsonByte, &obj)

	log.Println("Unmarshal err:", err)

	val, err := jsonpath.JsonPathLookup(obj, path)
	log.Println("jsonPathLookup :", val, err)

	json, err := json.Marshal(val)
	log.Println("marshal:", string(json))

	return
	l := &List{
		Id:   "123456",
		Data: make([]Item, 0),
	}
	l.Data = append(l.Data, Item{Age: 20, Salary: 200000})
	l.Data = append(l.Data, Item{Age: 30, Salary: 300000})

	UnL := &List{Data: make([]Item, 0)}

	fmt.Println("L:::", UnL)
	return

	i2 := test20()
	fmt.Println("i2 : ", i2)

	i2.Name = "kx update"
	fmt.Println("i2 update: ", i2)

	for _, v := range m {
		fmt.Println("old :", v)
	}
	return

	tm := make(map[string]string)
	limit := 3000000

	for j := 0; j < limit; j++ {
		key := strconv.Itoa(j)
		value := "zz_" + key
		tm[key] = value
	}

	var count int32
	var wg sync.WaitGroup
	//var mu sync.RWMutex

	for j := 0; j < 100000; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			//mu.RLock()
			//defer mu.RUnlock()
			key := strconv.Itoa(j)
			val := tm[key]
			fmt.Println("val:", val)
			atomic.AddInt32(&count, 1)
		}(j)
	}
	wg.Wait()
	fmt.Println("over:", count)
	return

	var t int

	//4311 7995   4298 6411
	for i := 0; i < 30000; i++ {
		t += i
	}
	fmt.Println("t:", t) // 449985000   4288 9532
	return

	i := Info{Name: "zz"}
	fmt.Println(i)
	i.Lock()
	defer i.Unlock()
	return
	no := 2
	num := fmt.Sprintf("%02d", no)
	fmt.Println("res:", num)

	return
	info := &Info{}
	fmt.Println(err, info)
	return

	//todo-gcflags=-m参数 查看具体堆栈情况
	//test15()

	start := time.Date(int(1970), time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(int(1970), time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(start, end)

	str := "ABc"
	fmt.Println(test16(str))
}
func test16(s string) string {
	var buffer bytes.Buffer
	for idx, chr := range s {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if idx > 0 {
				buffer.WriteRune('_')
			}
			chr -= ('A' - 'a')
		}
		buffer.WriteRune(chr)
	}

	return strings.ToLower(buffer.String())
}

func test15() {
	dir, _ := os.Getwd()
	fmt.Println("dir", dir)

	return
	start := time.Date(int(1970), time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(int(1970), time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(start, end)
}

func test14() {

	dir, _ := os.Getwd()
	file := dir + "/learn/00test/a/test.go.t"
	targetFile := dir + "/learn/00test/a/test.go"
	t, err := template.ParseFiles(file)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	f, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file err:%v", err)
		return
	}
	config := Info{Name: "zz"}
	err = t.Execute(f, config)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Println("**********ok*************")
}
func test13() {

	fmt.Println("📂 Do you want to override the folder ? ")

	return
	url := "https://www.baidu.com/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err:", err)
	}
	f, err := os.Create("./body2.txt")
	fmt.Println("err2:", err)
	io.Copy(f, resp.Body)
	return
	ei := []int{1, 2}
	ev := reflect.ValueOf(ei)
	fmt.Printf("ev类型=%v", ev.Kind())

	typ := reflect.TypeOf(ei)

	v := 2

	//todo typ 某个变量的反射类型，reflect.New(typ) 生成一个此类型的变量指针
	resultValue := reflect.New(typ).Elem()
	resultValue = reflect.Append(resultValue, reflect.ValueOf(v))
	fmt.Println("rv", resultValue.Interface())
	fmt.Printf("rv:%T,rv:%T", resultValue, resultValue.Interface())
	return

	n := 17
	n |= n >> 1 //8
	fmt.Println("n", n, 17^8, 1<<31, 1<<4, 2<<1, 2>>1, 2>>2)
}

const prime32 = uint32(16777619)

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func test12() {

	hash := uint32(2)

	hash *= prime32
	fmt.Println("hash", hash, uint32("a"[0]))
	return
	k := 19
	fmt.Printf("二进制=%b\n", k) //00010011 1*2^0 + 1*2^1 + 0*2^2 +  0*2^3 + 1*2^4 = 1+2+0+0+16=19

	fmt.Println("ddd", 19/2)
	fmt.Println("asss", 19<<2)
	//将一个数左移N位相当于将一个数乘以2^N，而将一个数右移N位相当于将这个数除以2^N。
	// 00010011=>00001001  右移1位
	k2 := k >> 1
	fmt.Println("19:", k, k2)
	fmt.Printf("二进制=%b\n", k2) //0000 1001  1*2^0 + 0*2^2 +  0*2^3 + 1*2^3 = 1+0+0+8=9

	return
	param := 20
	var n int
	if param <= 16 {
		n = 16
	}
	n = param - 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	if n < 0 {
		n = math.MaxInt32
	}
	n = n + 1
	fmt.Println("n", n)
}
func test11() {
	line := []byte{'a', 'b'}
	line = bytes.TrimSuffix(line, []byte{'b'})
	fmt.Printf("byte=%v", string(line))
}

func test10() {
	limit := 10
	res := testBackTrack(limit)
	fmt.Printf("最终结果=%d", res)
}

// todo 递归的执行流程和栈一样的，都是后进先出
/**
  运动开始时，首先为递归调用建立一个工作栈，其结构包括值参、局部变量和返回地址；

  每次执行递归调用之前，把递归函数的值参、局部变量的当前值以及调用后的返回地址压栈；

  每次递归调用结束后，将栈顶元素出栈，使相应的值参和局部变量恢复为调用前的值，然后转向返回地址指定的位置继续执行。
*/
func testBackTrack(i int) int {
	fmt.Printf("刚进来的i=%d\n", i)
	if i >= 13 {
		return i
	}
	testBackTrack(i + 1)

	//todo 后进先出
	//递归结束i=12
	//递归结束i=11
	//递归结束i=10
	fmt.Printf("递归结束i=%d\n", i)
	return i
}

func test2() int {
	res := sort.Search(100, func(i int) bool {
		return i >= 20
	})
	return res
}

func test() {
	res := Code(1) //类型转换
	s := Msg("test msg")
	fmt.Println(res, s)
	fmt.Printf("type %T", res)

	t := func(data int32) {
		fmt.Println(data)
	}

	t(int32(res))
}

func test03() []int {
	a := []int{1, 2, 3, 4}
	b := make([]int, len(a))
	c := a[0:4]
	copy(b, a)
	_ = c
	//fmt.Printf("%p %p %p \n", a, b, c)
	//fmt.Println(a, b, c)
	//fmt.Println(&a[0], &b[0], &c[0])
	return b
}

type C struct {
	Name string
}

func test04() *C {
	cccccccccccccc := &C{
		Name: "test",
	}
	return cccccccccccccc
}

func test05() {
	d := time.Second * 2
	for {
		t := time.NewTimer(d)
		<-t.C
		fmt.Println("time:", time.Now().Format("2006-01-02 15:04:05"))
	}
}

type User struct {
	Name string
	Age  int
}

func (u *User) ShowName(value interface{}) {
	fmt.Printf("User的名字是=%s\n", u.Name)
}

func test06() {
	//dest := 1
	//dest := []int{1, 2}
	dest := &User{
		Name: "zz",
		Age:  20,
	}
	r := reflect.ValueOf(&dest)
	destValue := reflect.Indirect(r)
	//fmt.Printf("r的值=%+v,destValue的值=%+v\n", r, destValue)
	//fmt.Printf("r的类型=%T,destValue的类型=%T\n", r, destValue)
	fmt.Println(destValue.FieldByName("Name").Interface())
}

type TF func(str string) string

func tf01(str string) string {
	return "tf01" + str
}

func test07(tf TF) string {
	tstr := "abc"
	return tf(tstr)
}

func test08() {
	res := digui(0)
	fmt.Println(res)
}

func digui(i int) string {
	if i == 5 {
		return "退出了" + strconv.Itoa(i)
	}
	i++
	return digui(i)
}

func test09() {

	data := make(map[int]*User)
	dest := make([]*User, 0)
	data[0] = &User{Name: "zz", Age: 18}
	data[1] = &User{Name: "kx", Age: 16}
	for i := 0; i < 4; i += 2 {
		itme, ok := data[i]
		if ok {
			itme.Name = "被更新了"
		} else {
			itme = &User{"不在", 30}
			dest = append(dest, itme)
		}
	}
	for _, item := range dest {
		fmt.Println("item", item)
	}
	fmt.Println(data)
}
