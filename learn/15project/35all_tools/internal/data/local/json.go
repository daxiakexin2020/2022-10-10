package local

import (
	"35all_tools/internal/model"
	"encoding/json"
	"errors"
	"log"
	"reflect"
)

type JsonRepository struct{}

var _ (model.JsonRepo) = (*JsonRepository)(nil)

func NewJsonRepository() model.JsonRepo {
	return &JsonRepository{}
}

func (jr *JsonRepository) JsonCheck(model *model.JsonCheck) (interface{}, error) {
	if _, err := jr.jsonCheck(model.Str); err != nil {
		return nil, err
	}
	return model.Str, nil
}

func (jr *JsonRepository) JsonToGolangStruct(model *model.JsonToGolangStruct) (interface{}, error) {

	data := model.Str
	dest, err := jr.jsonCheck(data)
	if err != nil {
		return nil, err
	}
	ret := "type " + model.Name + " struct " + "{"

	ret += ("\n")
	recursion := jr.recursion(ret, dest, false)
	recursion += "}"
	return recursion, nil
}

func (jr *JsonRepository) recursion(s string, data interface{}, addFlag bool) string {

	dest, ok := data.(map[string]interface{})
	if ok {
		for k, v := range dest {
			s += "	" + k + "	"
			valueType := jr.valueType(v)
			if valueType == "map" {
				tmp := " struct " + "{\n"
				valueType = jr.recursion(tmp, v.(map[string]interface{}), true)
			}
			if valueType == "slice" {
				tmp := "	[]"
				valueType = jr.recursion(tmp, v.([]interface{}), false)
			}
			s += (valueType)
			s += ("\n")
		}
	} else {
		valueType := jr.valueType(data)
		if valueType == "slice" {
			sdata, ok := data.([]interface{})
			if ok {
				log.Println("sdata:", sdata)
				rs1, ok1 := sdata[0].([]map[string]interface{})
				rs2, ok2 := sdata[0].(map[string]interface{})
				rs3, ok3 := sdata[0].(interface{})
				if ok1 && len(rs1) > 0 {
					tmp := " struct " + "{\n"
					valueType = jr.recursion(tmp, rs1[0], true)
				} else if ok2 {
					tmp := " struct " + "{\n"
					valueType = jr.recursion(tmp, rs2, true)
				} else if ok3 {
					valueType = jr.valueType(rs3)
				}
			}
		}
		s += (valueType)
		s += ("\n")
	}

	if addFlag {
		s += "}\n"
	}
	return s
}

func (jr *JsonRepository) valueType(v interface{}) string {
	rtype := reflect.ValueOf(v)
	log.Println("rtype.Kind()", rtype.Kind())
	switch rtype.Kind() {
	case reflect.Uint:
		return "uint"
	case reflect.Uint32:
		return "uint32"
	case reflect.Int:
		return "int"
	case reflect.Int8:
		return "int8"
	case reflect.Int16:
		return "int16"
	case reflect.Int32:
		return "int32"
	case reflect.Int64:
		return "int64"
	case reflect.Float32:
		return "float32"
	case reflect.Float64:
		return "float64"
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "string"
	case reflect.Map:
		//遍历，递归
		return "map"
	case reflect.Slice:
		//拿一个值，递归判断类型
		return "slice"
	default:
		return "un"
	}
}

func (jr *JsonRepository) jsonCheck(data interface{}) (map[string]interface{}, error) {
	ma, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	dest := map[string]interface{}{}
	if err = json.Unmarshal(ma, &dest); err != nil {
		return nil, errors.New("not json")
	}
	return dest, nil
}
