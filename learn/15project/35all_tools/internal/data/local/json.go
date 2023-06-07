package local

import (
	"35all_tools/internal/model"
	"encoding/json"
	"errors"
	"log"
	"strings"
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

	builder := strings.Builder{}
	builder.WriteString(ret)
	builder.WriteString("\n")

	for k, _ := range dest {
		builder.WriteString(k)
		builder.WriteString("\n")
	}
	log.Println("换:", builder.String())
	return builder.String(), nil
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
