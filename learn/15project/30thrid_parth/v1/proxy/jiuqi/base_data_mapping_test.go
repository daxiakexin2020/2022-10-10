package jiuqi

import (
	"encoding/json"
	"fmt"
	"testing"
)

type BaseDataMappingTest struct {
	Code string                 `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func TestNewBaseDataMapping(t *testing.T) {

	data := map[string]interface{}{
		"Id":        "1",
		"TableName": "test_teable_name",
	}
	b := &BaseDataMappingTest{
		Code: "0",
		Msg:  "ok",
		Data: data,
	}
	mapping := NewBaseDataMapping()
	marshal, _ := json.Marshal(b)
	json.Unmarshal(marshal, &mapping)
	fmt.Println("data:::", mapping.Data)
}
