package service

import (
	"fmt"
	"reflect"
	"testing"
)

/**
reflect包实现了运行时反射，允许程序操作任意类型的对象。典型用法是用静态类型interface{}保存一个值，
调用TypeOf获取其动态类型信息，该函数返回一个Type类型值。
调用ValueOf函数返回一个Value类型值，该值代表运行时的数据。
Zero接受一个Type类型参数并返回一个代表该类型零值的Value类型值。
参见"The Laws of Reflection"获取go反射的介绍：http://golang.org/doc/articles/laws_of_reflection.html
*/

type User struct {
	Name   string
	Age    int
	salary uint
}

var (
	tint    = 1
	tbool   = true
	tstruct = User{Name: "zz", Age: 18, salary: 20}
	tmap    = make(map[string]string)
	tslice  = make([]int, 0)
)

func TestType_Of(t *testing.T) {

	fmt.Printf("【int 反射TypeOf=%v,类型=%T】\n", testTypeOf(tint), testTypeOf(tint)) //int，*reflect.rtype

	fmt.Printf("【bool 反射TypeOf=%v,类型=%T】\n", testTypeOf(tbool), testTypeOf(tbool)) //bool，*reflect.rtype

	fmt.Printf("【struct 反射TypeOf=%v】\n", testTypeOf(tstruct))

	fmt.Printf("【map 反射TypeOf=%v】\n", testTypeOf(tmap))

	fmt.Printf("【slice 反射TypeOf=%v】\n", testTypeOf(tslice))

}

func TestValue_Of(t *testing.T) {

	fmt.Printf("【int 反射ValueOf=%v，类型=%T】\n", testValueOf(tint), testValueOf(tint)) //1, reflect.Value

	fmt.Printf("【bool 反射ValueOf=%v】\n", testValueOf(tbool))

	fmt.Printf("【struct 反射ValueOf=%v】\n", testValueOf(tstruct))

	fmt.Printf("【map 反射ValueOf=%v】\n", testValueOf(tmap))

	fmt.Printf("【slice 反射ValueOf=%v】\n", testValueOf(tslice))
}

func Test_New(t *testing.T) {
	sliceType := reflect.TypeOf(tslice)
	newValue := testNew(sliceType)
	a := newValue.Interface().(*[]int)
	fmt.Printf("切片的新的Value=%v，类型是%T\n", a, a)
}

func Test_Indirect(t *testing.T) {
	tv := testValueOf(tstruct)
	res := testIndirect(tv)
	fmt.Printf("indirect Value=%v,类型=%T\n", res, res)
	fmt.Printf("indirect Value Type=%v,类型=%T\n", res.Type(), res.Type().Name())
}

func TestFUNC_NumField(t *testing.T) {
	tv := testValueOf(tstruct)
	tvType := testIndirect(tv).Type()
	for j := 0; j < tvType.NumField(); j++ {
		field := tvType.Field(j)
		fmt.Println("结构体的字段信息", field, tv.FieldByName("Age").Interface())
	}
}

// 返回持有v持有的指针指向的值的Value。如果v持有nil指针，会返回Value零值；如果v不持有指针，会返回v。
func testIndirect(v reflect.Value) reflect.Value {
	return reflect.Indirect(v)
}

// TypeOf返回接口中保存的值的类型，TypeOf(nil)会返回nil。
func testTypeOf(value interface{}) reflect.Type {
	return reflect.TypeOf(value)
}

// ValueOf返回一个初始化为i接口保管的具体值的Value，ValueOf(nil)返回Value零值。
func testValueOf(value interface{}) reflect.Value {
	return reflect.ValueOf(value)
}

/*
*
New返回一个Value类型值，该值持有一个指向类型为typ的新申请的零值的指针，返回值的Type为PtrTo(typ)。
ei := []int{1, 2}

	ev := reflect.ValueOf(ei)
	fmt.Printf("ev类型=%v", ev.Kind())

	typ := reflect.TypeOf(ei)

	v := 2

	//todo typ 某个变量的反射类型，reflect.New(typ) 生成一个此类型的变量指针
	resultValue := reflect.New(typ).Elem()
	resultValue = reflect.Append(resultValue, reflect.ValueOf(v)) todo
	fmt.Println("rv", resultValue.Interface())
	fmt.Printf("rv:%T,rv:%T", resultValue, resultValue.Interface())
	return

	常用来生成一个指定类型的变量
*/
func testNew(typ reflect.Type) reflect.Value {
	return reflect.New(typ)
}
