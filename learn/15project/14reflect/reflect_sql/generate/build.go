package generate

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Build struct {
	sql  string
	vals []interface{}
}

func B(dest interface{}) (*Build, error) {
	destValue := reflect.ValueOf(dest)
	destType := reflect.TypeOf(dest)
	switch destType.Kind() {
	case reflect.Struct:
		fmt.Println("typo ok", destType.Name())
	case reflect.Ptr:
		fmt.Println("type ptr", destType.Name())
		destType = destType.Elem()
		destValue = destValue.Elem()
		if destType.Kind() != reflect.Struct {
			return nil, errors.New("ptr type inner error")
		}
	default:
		fmt.Println("type error", destType.Kind())
		return nil, errors.New("type error " + destType.Kind().String())
	}
	return b(destType, destValue)
}

func b(destType reflect.Type, destValue reflect.Value) (*Build, error) {
	fields := make([]string, 0)
	values := make([]interface{}, 0)
	for i := 0; i < destValue.NumField(); i++ {
		val := destValue.Field(i).Interface()
		tag, ok := destType.Field(i).Tag.Lookup("zorm")
		if ok {
			fields = append(fields, tag)
			values = append(values, val)
		}
	}
	fstr := strings.Join(fields, ",")
	sql := fmt.Sprintf("INSERT into %s (%s) VALUES", destType.Name(), fstr)

	nb := &Build{
		sql:  sql,
		vals: values,
	}
	fmt.Println("sql:", sql, nb)
	return nb, nil
}
