package generate

import (
	"errors"
	"fmt"
	"reflect"
)

type Build struct {
}

func B(dest interface{}) error {
	destValue := reflect.ValueOf(dest)
	destType := reflect.TypeOf(dest)
	switch destType.Kind() {
	case reflect.Struct:
		fmt.Println("类型正确", destType.Name())
	default:
		fmt.Println("类型错误，不能转换", destType.Kind())
		return errors.New("类型错误，不能转换")
	}
	for i := 0; i < destValue.NumField(); i++ {
		fmt.Println("k,v:", destType.Field(i).Tag.Get("zorm"), destValue.Field(i))
	}
	return nil
}
