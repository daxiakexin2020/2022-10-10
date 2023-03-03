package main

import (
	"fmt"
	"reflect"
)

type P struct {
	Name   string `json:"name"`
	Age    int
	salary int
	F      func() error
}

func (p *P) ShowName() string {
	return p.Name
}

func (p *P) ShowAge() int {
	return p.Age
}

func (p *P) showSalary() int {
	return p.salary
}

func main() {
	//ReflectType()
	ReflectValue()
}

/*
*
Value为go值提供了反射接口。

不是所有go类型值的Value表示都能使用所有方法。请参见每个方法的文档获取使用限制。在调用有分类限定的方法时，应先使用Kind方法获知该值的分类。调用该分类不支持的方法会导致运行时的panic。

Value类型的零值表示不持有某个值。零值的IsValid方法返回false，其Kind方法返回Invalid，而String方法返回"<invalid Value>"，所有其它方法都会panic。
绝大多数函数和方法都永远不返回Value零值。如果某个函数/方法返回了非法的Value，它的文档必须显式的说明具体情况。

如果某个go类型值可以安全的用于多线程并发操作，它的Value表示也可以安全的用于并发。
*/
func ReflectValue() {

	//反射value后，可以有一些方法，直接可以拿到value持有的inteface的值，比如  Int() String()...

	//valueType := &P{Name: "zz"}
	//valueType := "abc"
	valueType := make(map[string]int)
	valueType["a"] = 1

	reflectValue := reflect.ValueOf(valueType)
	fmt.Printf("ValueOf=%v\n", reflectValue)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	//IsValid返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。绝大多数函数和方法都永远不返回Value零值。如果某个函数/方法返回了非法的Value，它的文档必须显式的说明具体情况。
	valid := reflectValue.IsValid()
	fmt.Printf("vaild=%v\n", valid)

	//Kind返回v持有的值的分类，如果v是Value零值，返回值为Invalid
	kind := reflectValue.Kind()
	fmt.Printf("Kind()=%v\n", kind)

	//返回v持有的值的字符串表示。因为go的String方法的惯例，Value的String方法比较特别。
	//和其他获取v持有值的方法不同：v的Kind是String时，返回该字符串；v的Kind不是String时也不会panic而是返回格式为"<T value>"的字符串，其中T是v持有值的类型。
	s := reflectValue.String()
	fmt.Printf("String()=%v\n", s)

	if reflectValue.Kind() == reflect.Map {
		keys := reflectValue.MapKeys()
		fmt.Printf("MapKeys()=%v\n", keys)
	}

	//如果v持有的值可以被修改，CanSet就会返回真。只有一个Value持有值可以被寻址同时又不是来自非导出字段时，它才可以被修改。如果CanSet返回假，调用Set或任何限定类型的设置函数（如SetBool、SetInt64）都会panic。
	set := reflectValue.CanSet()
	fmt.Printf("CanSet()=%v\n", set)

	//反射修改原始值，先判断是否可以被修改，然后需要继续判断内部属性是否可以被修改，例如结构体，本身可以被修改，但是内部的属性，有的是小写，不可以导出，所以有的不可以修改
	if reflectValue.CanSet() {
		if reflectValue.Kind() == reflect.Struct {
			for i := 0; i < reflectValue.NumField(); i++ {
				//fmt.Printf("reflectValue.Field(i)类型=%T\n", reflectValue.Field(i))
				fmt.Println("val:", reflectValue.Field(i).SetString)
				if reflectValue.Field(i).CanSet() {
					if reflectValue.Field(i).Kind() == reflect.String {
						idata := reflectValue.Field(i).Interface()
						fmt.Printf("idata类型=%T\n", idata)
						if data, ok := idata.(string); ok {
							fmt.Printf("data=%v,type=%T\n", data, data)
						}
						reflectValue.Field(i).SetString("kx")
					}
					if reflectValue.Field(i).Kind() == reflect.Int {
						reflectValue.Field(i).SetInt(200)
					}
				}
			}
		}
	}

	fmt.Println("rdata:", valueType)

	//New() New返回一个Value类型值，该值持有一个指向类型为typ的新申请的零值的指针，返回值的Type为PtrTo(typ)。
	//通过反射生成新的变量
	newType := reflect.New(reflect.TypeOf(valueType))
	newType.Elem().Set(reflect.MakeMap(reflect.TypeOf(valueType)))

	if iidata, ok := newType.Interface().(*map[string]int); ok {
		fmt.Println("断言成功", iidata)
		(*iidata)["2"] = 2
		(*iidata)["3"] = 3

		TestReflectMap(*iidata)
	}
	fmt.Println("newData:", newType.Interface())
	fmt.Printf(" newType.Interface()类型=T\n", newType.Interface())
}

func TestReflectMap(data map[string]int) {
	fmt.Println("data:", data)
}

/*
*
Type类型用来表示一个go类型。
不是所有go类型的Type值都能使用所有方法。请参见每个方法的文档获取使用限制。
在调用有分类限定的方法时，应先使用Kind方法获知类型的分类。调用该分类不支持的方法会导致运行时的panic。
*/
func ReflectType() {
	//var refType *int
	//refType := make([]int, 10)
	//refType := [3]int{11, 11, 11}

	refType := &P{}

	//返回接口中保存的指的类型
	rTYPE := reflect.TypeOf(refType)
	fmt.Printf("typeof=%v\n", rTYPE) //[]int  类型

	//反射的值的类型的种类，  一共分为2种类型，一种是值的类型，另外一种是值的类型所属的种类 （Kind返回该接口的具体分类）
	fmt.Printf("typeof=%v\n", rTYPE.Kind()) //silce 具体的类型

	//Name返回该类型在自身包内的类型名，如果是未命名类型会返回"",比如反射的是结构体，就比较有意义
	fmt.Printf("Name()=%s\n", rTYPE.Name())

	//返回array类型的长度，如非数组类型将panic
	//fmt.Printf("Len()=%d\n", rTYPE.Len())

	//返回结构体类型的字段数量，如果不是struct，会panic ，如果是传的地址，则应该使用  rTYPE.Elem().NumField()
	if rTYPE.Kind() == reflect.Ptr {
		fmt.Println("rTYPE.Elem()", rTYPE.Elem().Kind() == reflect.Struct)
		fmt.Printf("NumField=%d\n", rTYPE.Elem().NumField())
	} else {
		fmt.Printf("NumField=%d\n", rTYPE.NumField())
	}

	if rTYPE.Kind() == reflect.Ptr {
		if rTYPE.Elem().Kind() == reflect.Struct {
			field := rTYPE.Elem().Field(0)
			fmt.Printf("field.Name=%s,filed.Type=%v,field.Tag=%v,field.Tag.Json=%v\n", field.Name, field.Type, field.Tag, field.Tag.Get("json"))
		}
	} else {
		if rTYPE.Kind() == reflect.Struct {
			field := rTYPE.Field(0)
			fmt.Printf("field.Name=%s,filed.Type=%v,field.Tag=%v,field.Tag.Json=%v\n", field.Name, field.Type, field.Tag, field.Tag.Get("json"))
		}
	}
	if rTYPE.Kind() == reflect.Ptr {
		if rTYPE.Elem().Kind() == reflect.Struct {
			fmt.Printf("NumMethod ELem=%d\n", rTYPE.Elem().NumMethod())
		}
	} else {
		if rTYPE.Kind() == reflect.Struct {
			fmt.Printf("NumMethod=%d\n", rTYPE.NumMethod())
		}
	}

	// 返回该类型的元素类型，如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
	fmt.Printf("Elem()=%v\n", rTYPE.Elem().Kind())
}
